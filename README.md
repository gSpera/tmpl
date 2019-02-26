tmpl
====

tmpl(short for template) is a simple tool made in Go which combines a template file and data in a single file.

tmpl tries to be simple, no recursive path support, neither the possibility to use multiple files at once,
this because tmpl is mainly aimed to scripts and tools like Make

Syntax
======

`$ tmpl -template=template.tmpl -data=data.json -output=output.html`

This is it.

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

Sanitizing
==========

By default tmpl sanitizes HTML/JS/CSS and other types as defined in html/template (documentation)[https://golang.org/pkg/html/template],
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

Formatters
==========

tmpl uses external packages for unmarshaling data, by default it uses

|Format | Package | Registered as|
|-------|---------|--------------|
|json|encoding/json|json|
|yaml|gopkg.in/yaml.v2|yaml, yml|

You can read the respective licences at the project pages.

Contributing
============

Contribution are welcome, try to keep code simple and to not bloat too much the software.
