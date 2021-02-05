package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

func main() {
	// Routes
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", about)

	// 	Server
	http.ListenAndServe(":80", nil)
}
