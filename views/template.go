package views

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTemp *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.htmlTemp.Execute(w, data)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func Parse(fs fs.FS, tmplName ...string) Template {
	tmpl, err := template.ParseFS(fs, tmplName...)
	//if the gohtml template cannot be parsed, panic
	if err != nil {
		panic(err)
	}

	temp := Template{
		htmlTemp: tmpl,
	}

	return temp
}
