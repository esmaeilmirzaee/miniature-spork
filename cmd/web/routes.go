package main

import (
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/config"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux:=chi.NewMux()

	mux.Get("/", handlers.Repo.HomePageHandler)

	return mux
}
