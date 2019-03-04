package format

import (
	"encoding/json"
)

func init() {
	RegisterDataFormat("json", jsonUnmarshaler{})
}

type jsonUnmarshaler struct{}

func (jsonUnmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	err = json.Unmarshal(raw, &data)
	return
}
