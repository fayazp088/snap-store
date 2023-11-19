package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	fmt.Println("Raw path: ", r.URL.RawPath)
	fmt.Println("Route Path: ", r.URL.Path)

	switch path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page Not Found!!!", http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Hello Home!!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Contact Page !!</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul> <li>
    <b>Is there a free version?</b>

	 Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
</li> <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:support@snapstore.com">support@snapstore.com</a>
  </li>
</ul> `)

}

func main() {
	var router Router
	http.ListenAndServe(":3000", router)
}
