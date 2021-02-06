package main

import (
	"net/http"

	"github.com/NathanielRand/GoBnB/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	// Routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)

	// 	Server
	http.ListenAndServe(portNumber, nil)
}
