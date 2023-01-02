package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Conor-Fleming/lenslocked/controllers"
	"github.com/Conor-Fleming/lenslocked/templates"
	"github.com/Conor-Fleming/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tmpl := views.Parse(templates.FS, "home.gohtml", "contact.gohtml", "faq.gohtml")
	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl = views.Parse(templates.FS, "contact.gohtml")
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl = views.Parse(templates.FS, "faq.gohtml")
	r.Get("/faq", controllers.FAQ(tmpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("starting the server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
