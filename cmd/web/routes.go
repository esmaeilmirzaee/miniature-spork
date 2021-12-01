package main

import (
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/config"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux:=chi.NewMux()

	//Middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomePageHandler)
	mux.Get("/about", handlers.Repo.AboutPageHandler)

	return mux
}
