package main

import (
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/config"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/handlers"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/renderers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	app.UseCache = false
	tc, err := renderers.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	renderers.NewTemplate(&app)
	
	http.HandleFunc("/", handlers.Repo.HomePageHandler)
	log.Println("App is listening on " + portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}