package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, filepath.Join("templates", "home.gohtml"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, filepath.Join("templates", "contact.gohtml"))
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

type User struct {
	Name string
	Age  int
	Bio  string
}

func executeTemplate(w http.ResponseWriter, path string) {
	t, err := template.ParseFiles(path)

	if err != nil {
		log.Printf("Error while parsing teampleate")
		http.Error(w, "Error parsing template!!!", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Printf("Error while executing teampleate")
		http.Error(w, "Error executing template!!!", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	t, err := template.ParseFiles("./test.gohtml")

	if err != nil {
		panic(err)
	}

	user := User{
		Name: "fayaz pasha",
		Age:  27,
		Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":3000", r)
}
