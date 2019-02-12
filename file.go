package main

import (
	"fmt"
	"os"
)

//FileFlag can be used to specify the path to an existing file
type FileFlag string

//String returns a string rapresentation of the Flag
func (f *FileFlag) String() string {
	if f == nil {
		return "<invalid>"
	}

	return string(*f)
}

//Set sets the value to the flag, if the file doesn't exist Set will return an error
func (f *FileFlag) Set(value string) error {
	if f == nil {
		return fmt.Errorf("Cannot access nil FileFlag")
	}

	*f = FileFlag(value)
	fmt.Printf("FIleFlag.Set(%s) = %v\n", value, f)
	_, err := os.Stat(value)
	if err != nil { //This include if the file doesn't exist
		return err
	}

	return nil
}

//IsValid can be used to check if the value of the flag is invalid, this can be because of not initialization of the flag or by forcing the value
func (f *FileFlag) IsValid() bool {
	_, err := os.Stat(f.String())
	return err == nil
}
