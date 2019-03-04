package format

import (
	"bytes"
	"encoding/csv"
)

func init() {
	RegisterDataFormat("csv", csvUnmarshaler{})
}

type csvUnmarshaler struct{}

func (csvUnmarshaler) Unmarshal(raw []byte) (data interface{}, err error) {
	r := csv.NewReader(bytes.NewReader(raw))
	return r.ReadAll()
}
