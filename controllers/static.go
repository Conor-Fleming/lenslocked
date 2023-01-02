package controllers

import (
	"html/template"
	"net/http"

	"github.com/Conor-Fleming/lenslocked/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func FAQ(tmpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is this website about?",
			Answer:   "This website is about providing information and resources on a variety of topics.",
		},
		{
			Question: "How can I contact the website owner?",
			Answer:   "You can contact the website owner by sending an email to info@example.com.",
		},
		{
			Question: "Can I make suggestions for new content on the website?",
			Answer:   `Yes, you can send your suggestions to <a href="mailto:info@example.com">info@example.com</a>. We welcome feedback and new ideas for content.`,
		},
		{
			Question: "Is this website free to use?",
			Answer:   "Yes, this website is free to use for all visitors. Oh yeah.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, questions)
	}
}
