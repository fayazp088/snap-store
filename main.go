package main

import (
	"net/http"

	"github.com/fayazp088/snap-store/controllers"
	"github.com/fayazp088/snap-store/templates"
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

	homeTpl := views.Must(views.ParseFs(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))
	r.Get("/", controllers.StaticHandler(homeTpl))

	contactTpl := views.Must(views.ParseFs(
		templates.FS,
		"contact.gohtml", "tailwind.gohtml",
	))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	faqTpl := views.Must(views.ParseFs(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(faqTpl))

	//Users
	var userC controllers.Users

	userC.Template = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", userC.New)
	r.Post("/signup", userC.Create)

	loginTpl := views.Must(views.ParseFs(templates.FS, "login.gohtml", "tailwind.gohtml"))
	r.Get("/login", controllers.StaticHandler(loginTpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", r)
}
