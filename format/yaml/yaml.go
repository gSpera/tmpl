package json

import (
	"gopkg.in/yaml.v2"

	"github.com/gSpera/tmpl/format"
)

func init() {
	format.RegisterDataFormat("yaml", unmarshaler{})
	format.RegisterDataFormat("yml", unmarshaler{})
}

type unmarshaler struct{}

func (unmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	err = yaml.Unmarshal(raw, &data)
	return
}
