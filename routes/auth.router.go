package routes

import (
	"github.com/gorilla/mux"
	"github.com/morshedulmunna/pxomart-api/app/controllers"
	"github.com/morshedulmunna/pxomart-api/app/repositories"
	"github.com/morshedulmunna/pxomart-api/app/services"
)

// AuthRoutes registers authentication-related routes
func AuthRoutes(router *mux.Router) {
	// Initialize repository, service, and controller
	authRepo := repositories.NewAuthRepository()
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	// Register routes
	router.HandleFunc("/login", authController.Login).Methods("POST")
}
