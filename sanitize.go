package main

import "html/template"

//safeTemplateFunctions is a FuncMap that contains functions for inserting data without sanitization
var safeTemplateFunctions = template.FuncMap{
	"safeHTML": func(s string) template.HTML {
		return template.HTML(s)
	},
	"safeJS": func(s string) template.JS {
		return template.JS(s)
	},
	"safeCSS": func(s string) template.CSS {
		return template.CSS(s)
	},
	"safeURL": func(s string) template.URL {
		return template.URL(s)
	},
	"safeHTMLAttr": func(s string) template.HTMLAttr {
		return template.HTMLAttr(s)
	},
	"safeJSStr": func(s string) template.JSStr {
		return template.JSStr(s)
	},
	"safeSrcset": func(s string) template.Srcset {
		return template.Srcset(s)
	},
}
