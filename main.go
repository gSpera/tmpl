package main

import (
	"fmt"
	"plugin"

	_ "github.com/gSpera/tmpl/format/json"

	"github.com/gSpera/tmpl/format"
)

func main() {
	p, err := plugin.Open("yaml.so")
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded plugin:", p)
	fmt.Println("Registered: ", format.Registered())
}
