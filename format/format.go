package format

import (
	"strings"
)

//Format is a decoding format
type Format struct {
	Name        string
	Unmarshaler Unmarshaler
}

//InvalidUnmarshaler is used when no valid unmarshaler can be found
var InvalidUnmarshaler = Format{"INVALID", nil}

//Unmarshal use the Format unmarshaler to unmarshal the data
func (f *Format) Unmarshal(raw []byte) (interface{}, error) {
	return f.Unmarshaler.Unmarshal(raw)
}

var formats []Format

//Unmarshaler is used to unmarshal the data
type Unmarshaler interface {
	//Unmarshal takes the raw data as input and decodes it into the interface{}
	Unmarshal(raw []byte) (interface{}, error)
}

//RegisterDataFormat manages input format for data
func RegisterDataFormat(name string, unmarshaler Unmarshaler) {
	if formats == nil {
		formats = make([]Format, 0, 10)
	}
	formats = append(formats, Format{
		Name:        name,
		Unmarshaler: unmarshaler,
	})
}

//Registered returns a slice with all the registered formats
func Registered() []Format { return formats }

//FindBest searches for the best format for decoding the file
//This is done by searching the filename suffix(usually the extension) and the prefix
func FindBest(filename string) Format {
	for _, f := range formats {
		if strings.HasSuffix(filename, f.Name) || strings.HasPrefix(filename, f.Name) {
			return f
		}
	}

	return InvalidUnmarshaler
}

//ByName return the FOrmat with the given name
func ByName(name string) Format {
	for _, f := range formats {
		if f.Name == name {
			return f
		}
	}

	return InvalidUnmarshaler
}
