package format_test

import (
	"fmt"
	"testing"

	"github.com/gSpera/tmpl/format"
)

type badUnmarshaler struct{}

func (badUnmarshaler) Unmarshal([]byte) (interface{}, error) {
	return nil, fmt.Errorf("bad unmarshaler")
}

func TestRegisterDataFormat(t *testing.T) {
	u := badUnmarshaler{}
	format.RegisterDataFormat("format", u)
	for _, f := range format.Registered() {
		if f.Unmarshaler == u {
			return
		}
	}

	t.Error("registered format not found")
}

func TestUnmarshaler(t *testing.T) {
	u := badUnmarshaler{}
	format.RegisterDataFormat("format", u)

	f := format.ByName("format")
	if f == format.InvalidUnmarshaler {
		t.Fatal("cannot find the registered formatter")
	}

	_, _ = f.Unmarshal([]byte{})
}

func TestFindBest(t *testing.T) {
	u := format.Format{Name: "format", Unmarshaler: badUnmarshaler{}}
	tm := []struct {
		name     string
		filename string
		format   format.Format
	}{
		{
			"empty",
			"",
			format.InvalidUnmarshaler,
		},
		{
			"suffix",
			"a.format",
			u,
		},
		{
			"prefix",
			"format.a",
			u,
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			f := format.FindBest(tt.filename)
			if tt.format != f {
				t.Errorf("wrong formatter; expected: %v; got: %v", tt.format, f)
			}
		})
	}
}

func TestByName(t *testing.T) {
	u := format.Format{Name: "format", Unmarshaler: badUnmarshaler{}}
	tm := []struct {
		name     string
		filename string
		format   format.Format
	}{
		{
			"empty",
			"",
			format.InvalidUnmarshaler,
		},
		{
			"name",
			"format",
			u,
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			f := format.ByName(tt.filename)
			if tt.format != f {
				t.Errorf("wrong formatter; expected: %v; got: %v", tt.format, f)
			}
		})
	}
}
