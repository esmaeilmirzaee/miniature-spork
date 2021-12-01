package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/config"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/renderers"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// In production change InProduction to true
	app.InProduction = false

	// A new session manager configuration
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	app.UseCache = false
	tc, err := renderers.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	renderers.NewTemplate(&app)
	
	//http.HandleFunc("/", handlers.Repo.HomePageHandler)
	log.Println("App is listening on " + portNumber)
	//_ = http.ListenAndServe(portNumber, nil)

	srv := http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	srv.ListenAndServe()
}