package web

import (
	template2 "html/template"
	"os"
)

func htmlPath(name string) string {
	wd, _ := os.Getwd()
	return wd + "/pkg/template/" + name + ".html"
}

func template(name string) *template2.Template {
	return template2.Must(template2.ParseFiles(htmlPath(name)))
}
