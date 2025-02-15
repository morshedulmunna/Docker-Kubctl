package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// HealthRoutes sets up health-related routes
func WelcomeRoutes(version map[string]string) *chi.Mux {
	r := chi.NewRouter()

	// Root route, returns a general status or welcome message in JSON
	r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Extract version keys (v1, v2, v3) for the response
		availableVersions := make([]string, 0, len(version))
		for key := range version {
			availableVersions = append(availableVersions, version[key])
		}

		response := map[string]any{
			"message":            "Welcome to the Server",
			"available_versions": availableVersions,
		}

		json.NewEncoder(w).Encode(response)
	}))

	return r
}
