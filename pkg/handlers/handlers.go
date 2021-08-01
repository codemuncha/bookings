package handlers

import (
	"github.com/tsawler/bookings/pkg/config"
	"github.com/tsawler/bookings/pkg/models"
	"github.com/tsawler/bookings/pkg/render"
	"net/http"
)

//TemplateData holds data send from templates package

var Repo *Repository


type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repo
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)


	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository)  About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
