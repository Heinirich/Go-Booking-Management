package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/heinirich/bookings/pkg/config"
	"github.com/heinirich/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler{

	mux  := chi.NewRouter()

	mux.Use(writeToConsole)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	mux.Get("/home",handlers.Repo.HomePage)
	mux.Get("/about",handlers.Repo.AboutPage)

	return mux
}
