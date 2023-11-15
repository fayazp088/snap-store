package main

import (
	"fmt"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome Golang!!</h1>")
	
}

func main() {
	http.HandleFunc("/", healthCheck)

	http.ListenAndServe(":3000", nil)
}
