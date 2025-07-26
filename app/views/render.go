package render

import (
	"html/template"
	"io"
)

const templateDir = "app/views/"

var (
	index = parse(templateDir+"layout.html", templateDir+"index.html")
	about = parse(templateDir+"layout.html", templateDir+"about.html")
)

func Index(w io.Writer, data map[string]interface{}, partial string) error {
	if partial == "" {
		partial = "layout"
	}
	return index.ExecuteTemplate(w, partial, data)
}

func About(w io.Writer, data map[string]interface{}, partial string) error {
	if partial == "" {
		partial = "layout"
	}
	return about.ExecuteTemplate(w, partial, data)
}

func parse(files ...string) *template.Template {
	return template.Must(template.ParseFiles(files...))
}
