package main

import (
	"log"
	"plugin"
)

//registerExternalFunctions open the library at the given path
//this functions logs if an error occour, this will always appen on windows, becauyse Go does not support plugin on windows
func registerExternalFunction(path string) {
	_, err := plugin.Open(path)
	if err != nil {
		log.Printf("Cannot load %s: %v\n", path, err)
	}
}
