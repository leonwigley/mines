package render

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *
var files embed.FS

var (
	index = parse("index.html")
	about = parse("about.html")
)

func Index(w io.Writer, data map[string]interface{}, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return index.ExecuteTemplate(w, partial, data)
}

func About(w io.Writer, data map[string]interface{}, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return about.ExecuteTemplate(w, partial, data)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}
