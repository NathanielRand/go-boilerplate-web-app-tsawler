package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact")
}

func main() {
	// Routes
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", about)

	// 	Server
	http.ListenAndServe(":80", nil)
}
