package format

import (
	"github.com/BurntSushi/toml"
)

func init() {
	RegisterDataFormat("toml", tomlUnmarshaler{})
}

type tomlUnmarshaler struct{}

func (tomlUnmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	err = toml.Unmarshal(raw, &data)
	return
}
