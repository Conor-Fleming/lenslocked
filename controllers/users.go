package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Executer
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Parsing form:", http.StatusInternalServerError)
		return
	}

	email := r.PostForm.Get("email")
	pass := r.PostForm.Get("password")
	fmt.Fprintf(w, "<p>Email: %s</p>\n<p>Password: %s</p>", email, pass)
}
