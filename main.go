package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

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
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">jon@calhoun.io</a>.</p>")
}

func main() {
	r := chi.NewRouter
	r.Method("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("starting the server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
