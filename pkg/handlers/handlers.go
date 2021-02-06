package handlers

import (
	"net/http"

	"github.com/NathanielRand/GoBnB/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.html")
}
