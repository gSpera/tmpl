tmpl
====

tmpl(short for template) is a simple tool made in Go which combines a template file and data in a single file.

tmpl tries to be simple, no recursive path support, neither the possibility to use multiple files at once,
this because tmpl is mainly aimed to scripts and tools like Make

Syntax
======

`$ tmpl -template=template.tmpl -data=data.json -output=output.html`

That's it.

File type
=========

tmpl doesn't care about file type, you can use it for whatever you want, from generating HTML pages to simple text.

The template file may be provided with Go text/template syntax, data is a little more complicated, plugin can be loaded at runtime or can be embedded inside the executable, the extension of the data file is used to determine the format.

Options
=======

| Long name | Short Name | Description | 
|-----------|------------|-------------| 
| -template | -t | Template file | 
| -data | -d | Data file | 
| -output | -o | Output file | 
| -help | -h | Show help | 
| -functions| -fn | Functions file |

Sanitizing
==========

By default tmpl sanitizes HTML/JS/CSS and other types as defined in html/template [documentation](https://golang.org/pkg/html/template),
if you need to insert raw data inside the template you can use the following functions:

|Type | Function |
|-----|----------|
|HTML | safeHTML|
|JS | safeJS|
|CSS | safeCSS|
|URL | safeURL|
|HTMLAttr | safeHTMLAttr|
|JSStr | safeJSStr|
|Srcset | safeSrcset|

Functions
=========

Over sanitizing functions tmpl provides some simple functions that you can use inside a template

|Name |Params |Return |Usage |
|-----|-------|-------|------|
|currentTime | | time.Time | Returns the current time.Time, like calling time.Time in Go |
|parseTime |format(string), value(string) |time.Time | Parses the value string as time with the given format, like calling time.Parse in Go |
|fromUnixTime |unix(int64) |time.Time|Returns the time.Time at the given unix timestamp, like calling time.Unix(unix, 0) |
|fromUnixNanoTime |unix(int64), nano(int64) |time.Time |Returns the time.Time at the given unix timestamp, like calling time.Unix(unix, nano) |

Formatters
==========

tmpl uses external packages for unmarshaling data, by default it uses

|Format | Package | Registered as|
|-------|---------|--------------|
|json|encoding/json|json|
|csv|encoding/csv|csv|
|yaml|gopkg.in/yaml.v2|yaml, yml|
|toml|github.com/BurntSushi/toml|toml|

You can read the respective licences at the project pages.


Custom Formats and Functions
============================

tmpl searches a `tmpl.so` library in the workdir and loads it, this is usefull for registering custom functions and formats at runtime.
Use the `-externalFile` flag to specify a custom file

For creating a `tmpl.so` file the suggested way is to create a `tmpl.go` file and set it up with a main funcion(which, currently, is not called) and an init function that will register custom functions and formats, building is done by the following command:
```go build -o=tmpl.so -buildmode=plugin tmpl.go```

See also [Go plugin documentation](https://golang.org/pkg/plugin)
**This is not avaible on windows**


Contributing
============

Contribution are welcome, try to keep code simple and to not bloat too much the software.
