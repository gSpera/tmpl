package json

import (
	"encoding/xml"

	"github.com/gSpera/tmpl/format"
)

func init() {
	format.RegisterDataFormat("xml", unmarshaler{})
}

type unmarshaler struct{}

func (unmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	err = xml.Unmarshal(raw, &data)
	return
}
