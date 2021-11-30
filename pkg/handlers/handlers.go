package handlers

import (
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/renderers"
	"net/http"
)

// HomePageHandler handles index page requests.
func HomePageHandler(w http.ResponseWriter, r * http.Request) {
	renderers.RenderTemplate(w, "home.page.tmpl")
}