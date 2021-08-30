package main

import (
	"net/http"

	"github.com/GoldenGlow300/mytestapp/cmd/pkg/config"
	"github.com/GoldenGlow300/mytestapp/cmd/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//middleware allows you to process a request and perform some action on it
	//Ex. checking to see if the user is authenticated before responding to a request
	mux.Use(middleware.Recoverer)
	

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
