package handlers

import (
	"net/http"

	"github.com/acharnovich/hotel-bookings/pkg/config"
	"github.com/acharnovich/hotel-bookings/pkg/models"
	"github.com/acharnovich/hotel-bookings/pkg/render"
)

// repo used by handlers
var Repo *Repository
// repo type
type Repository struct {
	App *config.AppConfig

}
// creates a repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}


func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["Test"] = "Hello"


	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
