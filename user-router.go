package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) userRouters() *chi.Mux {
	// middlewares

	// add routes
	a.App.Router.Get("/", a.Controller.Home)

	// file server
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Router.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Router
}
