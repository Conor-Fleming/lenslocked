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

// Display new user sign up
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

// Takes post action and parses the data (email password) from the submitted signup form
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Parsing form:", http.StatusInternalServerError)
		return
	}

	email := r.PostForm.Get("email")
	photo := r.PostForm.Get("photo")
	pass := r.PostForm.Get("password")
	fmt.Fprintf(w, "<p>Email: %s</p>\n<p>Password: %s\n</p>", email, pass)
	fmt.Fprintf(w, "<p>Photo: %v</p>", photo)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {

	//validate email and password and then load User home page
}
