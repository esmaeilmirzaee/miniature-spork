package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Writo to console")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRf protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{HttpOnly: true, Path: "/", Secure: app.InProduction, SameSite: http.SameSiteLaxMode})

	return csrfHandler
}

// SessionLoad saves and loads session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}