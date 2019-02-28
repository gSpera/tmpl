package function

import "time"

//This file registers the builtins functions
//To register a custom function create a separate file

func init() {
	RegisterFunction("now", time.Now)
	RegisterFunction("parseTime", time.Parse)
}
