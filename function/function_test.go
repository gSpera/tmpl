package function

import (
	"reflect"
	"testing"
)

func TestRegisterFunction(t *testing.T) {
	functions = nil //Force the initialization of functions

	RegisterFunction("func", TestRegisterFunction)

	if functions == nil {
		t.Error("functions is nil")
	}
	fn, ok := functions["func"]
	if !ok {
		t.Error("cannot get registered function")
	}
	if fn == nil {
		t.Error("function pointer is nil")
	}
}

func TestFuncMap(t *testing.T) {
	if !reflect.DeepEqual(functions, FuncMap()) {
		t.Error("functions is different from FuncMap()")
	}
}
