package format

import (
	"reflect"
	"testing"
)

func TestTOMLUnmarshal(t *testing.T) {
	type outputMap = map[string]interface{}
	tm := []struct {
		name  string
		input []byte
		out   interface{}
		err   error
	}{
		{
			"empty",
			[]byte(""),
			outputMap{},
			nil,
		},
		{
			"string",
			[]byte(`a = "b"`),
			outputMap{"a": "b"},
			nil,
		},
		{
			"multiple value",
			[]byte("a = 3\nb = \"c\""),
			outputMap{"a": int64(3), "b": "c"},
			nil,
		},
		{
			"array",
			[]byte(`a = ["a", "b", "c"]`),
			outputMap{"a": []interface{}{"a", "b", "c"}},
			nil,
		},
	}

	u := tomlUnmarshaler{}
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
