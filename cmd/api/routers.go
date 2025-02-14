package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	//Health check
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthChecker)
	})

	return r
}
