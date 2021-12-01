package handlers

import (
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/config"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/models"
	"github.com/esmaeilmirzaee/random-time-sleeper/pkg/renderers"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// HomePageHandler handles index page requests.
func (e *Repository) HomePageHandler(w http.ResponseWriter, r * http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hi. Bonjour."

	renderers.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (e *Repository) AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}