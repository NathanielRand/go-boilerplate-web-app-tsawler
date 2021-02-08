package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf func
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// Change value of "Secure" on cookie object to "true" in production.
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
