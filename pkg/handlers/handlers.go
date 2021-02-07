package handlers

import (
	"net/http"

	"github.com/NathanielRand/GoBnB/pkg/render"
)

// Home function renders the home page.
func Home(w http.ResponseWriter, r *http.Request) {
	render.RendersTemplate(w, "home.page.html")
}

// About function renders the about page.
func About(w http.ResponseWriter, r *http.Request) {
	render.RendersTemplate(w, "about.page.html")
}
