package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, filepath.Join("templates", "home.gohtml"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, filepath.Join("templates", "contact.gohtml"))
}

// ExecuteTemplate takes a response writer and a string representing the template name
// This will handle parsing the template and executing to prevent repeating this code for each static page
func executeTemplate(w http.ResponseWriter, tplName string) {
	tpl, err := template.ParseFiles(tplName)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the tempalte.", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
