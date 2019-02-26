package json

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestUnmarshal(t *testing.T) {
	tm := []struct {
		name  string
		input []byte
		out   interface{}
		err   error
	}{
		{
			"empty",
			[]byte(""),
			[][]string(nil),
			nil,
		},
		{
			"string",
			[]byte(`a,b`),
			[][]string{[]string{"a", "b"}},
			nil,
		},
		{
			"multiple value",
			[]byte("a,3\nb,c"),
			[][]string{[]string{"a", "3"}, []string{"b", "c"}},
			nil,
		},
		{
			"header",
			[]byte("name,age\na,3\nb,7"),
			[][]string{[]string{"name", "age"}, []string{"a", "3"}, []string{"b", "7"}},
			nil,
		},
	}

	u := unmarshaler{}
	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			out, err := u.Unmarshal(tt.input)
			if tt.err != err {
				t.Errorf("wrong error; expected: %v; got: %v\n", tt.err, err)
			}

			if !reflect.DeepEqual(tt.out, out) {
				spew.Dump(tt.out)
				spew.Dump(out)
				t.Errorf("wrong output; expected: %T{%v}; got: %T{%v}\n", tt.out, tt.out, out, out)
			}
		})
	}
}
