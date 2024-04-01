package handler

import (
	"log"
	"net/http"

	"github.com/Sakthiel/bookings/models"
	"github.com/Sakthiel/bookings/pkgs/config"
	"github.com/Sakthiel/bookings/pkgs/render"
)

var Repo *Repository

type Repository struct {
	AppConfig *config.Appconfig
}

func NewRepo(app *config.Appconfig) *Repository {
	return &Repository{
		AppConfig: app,
	}
}

func NewHandler(repo *Repository) {
	Repo = repo
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	log.Println(remoteIP)
	m.AppConfig.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "hello from sakthi"

	remoteIP := m.AppConfig.Session.GetString(r.Context() , "remoteIP")
	stringMap["remoteIP"] = remoteIP
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
