package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	gogenz "github.com/morshedulmunna/gogenz-library/response"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	if db == nil {
		panic("Database connection is nil in SetupRoutes")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gogenz.SuccessResponse(w, 200, "Welcome to the pxomart Server")
	}).Methods("GET")

	// router.HandleFunc("/register", uc.RegisterHandler).Methods("POST")

	return router
}
