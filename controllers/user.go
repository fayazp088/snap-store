package controllers

import (
	"fmt"
	"net/http"

	"github.com/fayazp088/snap-store/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserName string
		Email    string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) SingIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserName string
		Email    string
	}

	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Unable to parse form submission.", http.StatusBadRequest)
		return
	}

	email, password := r.FormValue("email"), r.FormValue("password")
	user, err := u.UserService.Create(email, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User Created: %v", user)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Unable to parse form submission.", http.StatusBadRequest)
		return
	}

	email, password := r.FormValue("email"), r.FormValue("password")

	user, err := u.UserService.Authenticate(email, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid username or pasaword.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User Data: %v", user)
}
