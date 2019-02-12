package format

var formats []string

//RegisterDataFormat manages input format for data
func RegisterDataFormat(format string) {
	if formats == nil {
		formats = make([]string, 10)
	}
	formats = append(formats, format)
}

//Registered returns a slice with all the registered formats
func Registered() []string { return formats }
