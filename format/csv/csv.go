package json

import (
	"bytes"
	"encoding/csv"

	"github.com/gSpera/tmpl/format"
)

func init() {
	format.RegisterDataFormat("csv", unmarshaler{})
}

type unmarshaler struct{}

func (unmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	r := csv.NewReader(bytes.NewReader(raw))
	return r.ReadAll()
}
