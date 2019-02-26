package json

import (
	"github.com/BurntSushi/toml"

	"github.com/gSpera/tmpl/format"
)

func init() {
	format.RegisterDataFormat("toml", unmarshaler{})
}

type unmarshaler struct{}

func (unmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	err = toml.Unmarshal(raw, &data)
	return
}
