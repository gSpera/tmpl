package main

import (
	"github.com/gSpera/tmpl/function"
)

func init() {
	function.RegisterFunction("median", func(data []interface{}) float64 {
		values := make([]float64, len(data))
		for i, v := range data {
			values[i] = v.(float64)
		}

		var med float64
		for _, v := range values {
			med += v
		}
		med /= float64(len(values))
		return med
	})
}

func main() {}
