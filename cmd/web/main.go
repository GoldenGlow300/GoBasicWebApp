package main

import (
	"fmt"
	"log"
	"net/http"

	//Where handlers is the name of the folder
	"github.com/GoldenGlow300/mytestapp/cmd/pkg/config"
	"github.com/GoldenGlow300/mytestapp/cmd/pkg/handlers"
	"github.com/GoldenGlow300/mytestapp/cmd/pkg/render"
)

const PORT = ":8080"

//entry point of the app
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//Passes the Home Page in to be handle
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", PORT))

	_ = http.ListenAndServe(PORT, nil) //Listens on port 8080. The _ means if there's an error I done't care about it.
}

/*

go to localhost:8080 in web browser
then run go main.go

*/
