package handlers

import (
	"github.com/heinirich/bookings/pkg/config"
	"github.com/heinirich/bookings/pkg/models"
	"github.com/heinirich/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is a repository type
type Repository struct {
	App *config.AppConfig
}


// NewRepo Creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{App: a}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository){
	Repo = r
}

func(m *Repository) HomePage(w http.ResponseWriter,r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)
	render.RenderTemplate(w,"home.page.html",&models.TemplateData{})
}

func(m *Repository) AboutPage(w http.ResponseWriter,r *http.Request){
	// Get remote IP
	userIP := m.App.Session.GetString(r.Context(),"remote_ip")

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["Hello"] = "Hello, again"
	stringMap["remote_ip"] = userIP



	// Send some data to the template
	render.RenderTemplate(w,"about.page.html",&models.TemplateData{
		StringMap: stringMap,
	})

}


