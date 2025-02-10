package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/morshedulmunna/pxomart-api/app/services"
)

// LoginRequest struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthControllerInterface defines the contract for AuthController
type AuthControllerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}

// AuthController struct
type AuthController struct {
	authService services.AuthService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService services.AuthService) AuthControllerInterface {
	return &AuthController{authService: authService}
}

// Login handles user authentication
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
