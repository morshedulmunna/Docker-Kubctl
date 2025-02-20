package routes

import (
	"github.com/go-chi/chi"
	"github.com/morshedulmunna/pxomart-api/cmd/controllers"
	"github.com/morshedulmunna/pxomart-api/cmd/services"
)

// UserRoutes sets up user-related routes
func UserRoutes() *chi.Mux {
	r := chi.NewRouter()
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	r.Post("/", userController.CreateUser)
	r.Get("/{id}", userController.GetUser)
	r.Put("/{id}", userController.UpdateUser)
	r.Delete("/{id}", userController.DeleteUser)

	return r
}
