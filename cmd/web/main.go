package main

import (
	"fmt"
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
	// Change "UseCache" value to "true" in production.
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// Start web server.
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
