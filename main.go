package main

import (
	"net/http"

	"github.com/fayazp088/snap-store/controllers"
	"github.com/fayazp088/snap-store/models"
	"github.com/fayazp088/snap-store/templates"
	"github.com/fayazp088/snap-store/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

type User struct {
	Name string
	Age  int
	Bio  string
}

func main() {
	r := chi.NewRouter()
	cfg := models.DefaultPostgresConfig()

	db, err := models.Open(cfg)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false),
	)

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
	userC := controllers.Users{
		UserService: &userService,
	}

	userC.Templates.New = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", userC.New)
	r.Post("/signup", userC.Create)

	userC.Templates.SignIn = views.Must(views.ParseFs(templates.FS, "login.gohtml", "tailwind.gohtml"))
	r.Get("/login", userC.SingIn)
	r.Post("/login", userC.ProcessSignIn)

	r.Get("/user/me", userC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", csrfMw(r))
}
