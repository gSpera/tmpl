package json

import (
	"encoding/json"

	"github.com/gSpera/tmpl/format"
)

func init() {
	format.RegisterDataFormat("json", unmarshaler{})
}

type unmarshaler struct{}

func (unmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	err = json.Unmarshal(raw, &data)
	return
}
