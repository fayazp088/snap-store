package main

import (
	"net/http"
	"path/filepath"

	"github.com/fayazp088/snap-store/controllers"
	"github.com/fayazp088/snap-store/views"
	"github.com/go-chi/chi/v5"
)

type User struct {
	Name string
	Age  int
	Bio  string
}

func main() {
	r := chi.NewRouter()

	homeTpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(homeTpl))

	contactTpl := views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	faqTpl := views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(faqTpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", r)
}
