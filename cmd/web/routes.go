package main

import (
	"net/http"

	"github.com/Sakthiel/bookings/pkgs/config"
	"github.com/Sakthiel/bookings/pkgs/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.Appconfig) http.Handler{
	mux := chi.NewRouter()
	
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(SessionLoad)
	mux.Get("/" , handler.Repo.Home)
	mux.Get("/about" , handler.Repo.About)

	return mux

}