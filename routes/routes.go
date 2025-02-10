package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morshedulmunna/gogenz-library"
)

func SetupRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gogenz.SuccessResponse(w, 200, "Welcome to the pxomart Server")
	}).Methods("GET")

	// Routers
	AuthRoutes(router)

	router.Use(gogenz.LoggingMiddleware)
	return router
}
