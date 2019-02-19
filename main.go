package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"text/template"
)

const outputPermission = 0666

func main() {
	templateFile := new(FileFlag)
	dataFile := new(FileFlag)
	var outputFile string
	flag.Var(templateFile, "template", "Template file")
	flag.Var(dataFile, "data", "Data file")
	flag.StringVar(&outputFile, "output", "", "Output file")

	flag.Parse()

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

	dataBody, err := ioutil.ReadFile(dataFile.String())
	if err != nil {
		log.Fatalf("Cannot read file: %v\n", err)
	}
	var data interface{}
	err = json.Unmarshal(dataBody, &data)
	if err != nil {
		log.Fatalf("Cannot unmarshal data: %v\n", err)
	}

	output, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, outputPermission)
	if err != nil {
		log.Fatalf("Cannot open output file: %v\n", err)
	}

	t, err := template.ParseFiles(templateFile.String())
	if err != nil {
		log.Fatalf("Cannot parse template: %v\n", err)
	}

	err = t.Execute(output, data)
	if err != nil {
		log.Fatalf("Cannot execute template: %v\n", err)
	}
}
