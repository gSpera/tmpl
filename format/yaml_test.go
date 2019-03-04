package format

import (
	"reflect"
	"testing"
)

func TestYAMLUnmarshal(t *testing.T) {
	type outputMap = map[interface{}]interface{}
	tm := []struct {
		name  string
		input []byte
		out   interface{}
		err   error
	}{
		{
			"empty",
			[]byte(""),
			nil,
			nil,
		},
		{
			"string",
			[]byte(`{"a": "b"}`),
			outputMap{"a": "b"},
			nil,
		},
		{
			"multiple value",
			[]byte(`{"a": 3, "b": "c"}`),
			outputMap{"a": 3, "b": "c"},
			nil,
		},
		{
			"array",
			[]byte(`["a", "b", 3]`),
			[]interface{}{"a", "b", 3},
			nil,
		},
	}

	u := yamlUnmarshaler{}
	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			out, err := u.Unmarshal(tt.input)
			if tt.err != err {
				t.Errorf("wrong error; expected: %v; got: %v\n", tt.err, err)
			}

			if !reflect.DeepEqual(tt.out, out) {
				t.Errorf("wrong output; expected: %T{%v}; got: %T{%v}\n", tt.out, tt.out, out, out)
			}
		})
	}
}
