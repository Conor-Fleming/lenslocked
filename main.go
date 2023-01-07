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

	homeTmpl := views.Parse(templates.FS, "home.gohtml", "tailwind.gohtml")
	r.Get("/", controllers.StaticHandler(homeTmpl))

	contactTmpl := views.Parse(templates.FS, "contact.gohtml", "tailwind.gohtml")
	r.Get("/contact", controllers.StaticHandler(contactTmpl))

	faqTmpl := views.Parse(templates.FS, "faq.gohtml", "tailwind.gohtml")
	r.Get("/faq", controllers.FAQ(faqTmpl))

	var user controllers.Users
	user.Templates.New = views.Parse(templates.FS, "signup.gohtml", "tailwind.gohtml")
	r.Get("/users", user.New)
	r.Post("/users", user.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("starting the server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
