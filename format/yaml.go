package format

import (
	"gopkg.in/yaml.v2"
)

func init() {
	RegisterDataFormat("yaml", yamlUnmarshaler{})
	RegisterDataFormat("yml", yamlUnmarshaler{})
}

type yamlUnmarshaler struct{}

func (yamlUnmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	err = yaml.Unmarshal(raw, &data)
	return
}
