package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Conor-Fleming/lenslocked/controllers"
	"github.com/Conor-Fleming/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tpl := views.Parse(filepath.Join("templates", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Parse(filepath.Join("templates", "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Parse(filepath.Join("templates", "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("starting the server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
