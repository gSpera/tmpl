package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gSpera/tmpl/format"
	"github.com/gSpera/tmpl/function"

	"html/template"
)

const defaultOutputPermission = 0666

func main() {
	templateFile := new(FileFlag)
	dataFile := new(FileFlag)
	var outputFile string
	flag.Var(templateFile, "template", "Template file")
	flag.Var(dataFile, "data", "Data file")
	flag.StringVar(&outputFile, "output", "", "Output file")
	showFormats := flag.Bool("showFormats", false, "shows avaible formats and exits")
	outputPermission := flag.Int("permission", defaultOutputPermission, "permission used to create output file")
	selectedFormat := flag.String("format", "", "use the specified format, default to automatic")

	flag.Parse()

	if *showFormats {
		formats := format.Registered()
		for _, f := range formats {
			fmt.Println(f.Name)
		}
		return
	}

	if !templateFile.IsValid() || !dataFile.IsValid() || strings.TrimSpace(outputFile) == "" {
		if !templateFile.IsValid() {
			fmt.Fprintln(os.Stderr, "template does not exist", templateFile)
		}
		if !dataFile.IsValid() {
			fmt.Fprintln(os.Stderr, "data does not exist", dataFile)
		}
		if strings.TrimSpace(outputFile) == "" {
			fmt.Fprintln(os.Stderr, "output is not set", outputFile)
		}

		fmt.Fprintf(os.Stderr, "\nUsage: tmpl -template=template.tmpl -data=data.json -output=output.html\n")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	//Read Data
	var unmarshaler format.Format
	if *selectedFormat != "" {
		unmarshaler = format.ByName(*selectedFormat)
	} else {
		unmarshaler = format.FindBest(dataFile.String())
	}
	if unmarshaler == format.InvalidUnmarshaler {
		log.Fatalf("Cannot find a unmarshaler for the data: use -format=format to specify one, use -showFormats to see avaible formats")
	}

	dataBody, err := ioutil.ReadFile(dataFile.String())
	if err != nil {
		log.Fatalf("Cannot read file: %v\n", err)
	}
	data, err := unmarshaler.Unmarshal(dataBody)
	if err != nil {
		log.Fatalf("Cannot unmarshal data: %v\n", err)
	}

	//Create output file
	output, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(*outputPermission))
	if err != nil {
		log.Fatalf("Cannot open output file: %v\n", err)
	}

	//Create Template
	t := template.New("tmpl")
	t.Funcs(function.FuncMap())
	t.Funcs(safeTemplateFunctions)
	tmplBody, err := ioutil.ReadFile(templateFile.String())
	if err != nil {
		log.Fatalf("Cannot read template file: %v\n", err)
	}
	t, err = t.Parse(string(tmplBody))
	if err != nil {
		log.Fatalf("Cannot parse template: %v\n", err)
	}

	err = t.Execute(output, data)
	if err != nil {
		log.Fatalf("Cannot execute template: %v\n", err)
	}
}
