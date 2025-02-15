package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	routes "github.com/morshedulmunna/pxomart-api/cmd/routers"
)

var version = map[string]string{
	"v1": "/api/v1",
}

func (app *application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// Root Route
	r.Mount("/", routes.WelcomeRoutes(version))

	// Versioning
	r.Route(version["v1"], func(r chi.Router) {
		r.Mount("/health", routes.HealthRoutes())
		// user route Mounting
		r.Mount("/users", routes.UserRoutes())
	})

	// Handle Not Found
	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Route not found:", r.URL.Path) // Log the missing route
		w.Header().Set("Content-Type", "application/json")
		statusCode := http.StatusNotFound
		w.WriteHeader(statusCode)
		response := map[string]interface{}{
			"status_code":   statusCode,
			"error":         "Not Found",
			"message":       "The requested resource could not be found.",
			"requested_url": r.URL.Path,
		}
		json.NewEncoder(w).Encode(response)
	}))

	return r
}
