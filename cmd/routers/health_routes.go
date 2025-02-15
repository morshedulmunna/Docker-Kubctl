package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// HealthRoutes sets up health-related routes
func HealthRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Root route, returns a general status or welcome message in JSON
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "Welcome to the health service",
		}
		json.NewEncoder(w).Encode(response)
	})

	// Health check route, returns health status in JSON
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// Perform actual health check (e.g., DB connection check, etc.)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"status": "OK",
		}
		json.NewEncoder(w).Encode(response)
	})

	return r
}
