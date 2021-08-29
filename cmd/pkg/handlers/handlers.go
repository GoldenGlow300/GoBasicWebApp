//This controls how it will be called when imported (ex. handlers.Home)
package handlers

import (
	"net/http"

	"github.com/GoldenGlow300/mytestapp/cmd/pkg/render"
)

//Handles request and responses for the Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
