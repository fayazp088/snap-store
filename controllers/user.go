package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Template Template
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserName string
		Email    string
	}

	data.Email = r.FormValue("email")
	u.Template.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Unable to parse form submission.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "<p>Email: %s</p>", r.FormValue("email"))
	fmt.Fprintf(w, "<p>Password: %s</p>", r.FormValue("password"))
	fmt.Fprintf(w, "<p>Username: %s</p>", r.FormValue("username"))
}
