package function

import "html/template"

//Function is a function, this is used for documentation pourpose
type Function = interface{}

var functions template.FuncMap

//RegisterFunction registers a function
func RegisterFunction(name string, fn Function) {
	if functions == nil {
		functions = make(template.FuncMap)
	}
	functions[name] = fn
}

//FuncMap return the registered functions as template.FuncMap
//the returned value can be nil if there are no registered functions
func FuncMap() template.FuncMap {
	return functions
}
