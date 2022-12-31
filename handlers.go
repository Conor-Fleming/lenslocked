package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func galleryIDHandler(w http.ResponseWriter, r *http.Request) {
	galleryID := chi.URLParam(r, "galleryID")
	w.Write(([]byte(fmt.Sprintf("Gallery ID: %v", galleryID))))

}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<head>
		<title>FAQ Page</title>
	</head>
  	<body>
		<h1>Frequently Asked Questions</h1>
  
		<h2>What is this website about?</h2>
		<p>This website is about providing information and resources on a variety of topics.</p>
  
		<h2>How can I contact the website owner?</h2>
		<p>You can contact the website owner by sending an email to info@example.com.</p>
  
		<h2>Can I make suggestions for new content on the website?</h2>
		<p>Yes, you can send your suggestions to info@example.com. We welcome feedback and new ideas for content.</p>
  
		<h2>Is this website free to use?</h2>
		<p>Yes, this website is free to use for all visitors.</p>
  </body>`)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//system agnostic for funzies
	tplPath := filepath.Join("templates", "home.gohtml")

	tpl, err := template.ParseFiles(tplPath)
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

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//system agnostic for funzies
	tplPath := filepath.Join("templates", "contact.gohtml")

	tpl, err := template.ParseFiles(tplPath)
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
