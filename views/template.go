package views

import (
	"html/template"
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

func Parse(filepath string) Template {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		panic(err)
	}

	temp := Template{
		htmlTemp: tpl,
	}

	return temp
}

/*func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}*/
