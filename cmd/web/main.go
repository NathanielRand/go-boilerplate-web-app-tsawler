package main

import (
	"log"
	"net/http"

	"github.com/NathanielRand/GoBnB/pkg/config"
	"github.com/NathanielRand/GoBnB/pkg/handlers"
	"github.com/NathanielRand/GoBnB/pkg/render"
)

const portNumber = ":8080"

func main() {
	// Template cache
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	
	app.TemplateCache = tc
	// Change to true in production.
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// Routes
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	// Server
	http.ListenAndServe(portNumber, nil)
}
