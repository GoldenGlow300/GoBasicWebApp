//This controls how it will be called when imported (ex. handlers.Home)
package handlers

import (
	"net/http"

	"github.com/GoldenGlow300/mytestapp/cmd/pkg/config"

	"github.com/GoldenGlow300/mytestapp/cmd/pkg/models"
	"github.com/GoldenGlow300/mytestapp/cmd/pkg/render"
)

//using the repository pattern
//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Handles request and responses for the Home Page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Again"

	//send data to the template

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
