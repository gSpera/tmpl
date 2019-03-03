package function

import "time"

//This file registers the builtins functions
//To register a custom function create a separate file

func init() {
	RegisterFunction("currentTime", time.Now)
	RegisterFunction("parseTime", time.Parse)
	RegisterFunction("fromUnixTime", func(unix int64) time.Time { return time.Unix(unix, 0) })
	RegisterFunction("fromUnixNanoTime", time.Unix)
}
