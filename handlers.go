package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Conor-Fleming/lenslocked/views"
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
	tpl, err := views.Parse(tplName)
	if err != nil {
		log.Printf("Parsing: %v", err)
		http.Error(w, "There was an error when parsing the template", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
}
