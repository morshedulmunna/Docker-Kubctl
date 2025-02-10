package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/morshedulmunna/pxomart-api/app/services"
)

// LoginRequest struct with validation tags
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// AuthControllerInterface defines the contract for AuthController
type AuthControllerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}

// AuthController struct
type AuthController struct {
	authService services.AuthService
	validator   *validator.Validate
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService services.AuthService) AuthControllerInterface {
	return &AuthController{
		authService: authService,
		validator:   validator.New(),
	}
}

// Login handles user authentication
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate request
	if validationErrs := c.validateRequest(req); validationErrs != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"errors": validationErrs})
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

// validateRequest performs validation on incoming request data
func (c *AuthController) validateRequest(req LoginRequest) map[string]string {
	errors := make(map[string]string)
	err := c.validator.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			errors[fieldName] = c.getCustomMessage(fieldName, err.Tag())
		}
		return errors
	}
	return nil
}

// getCustomMessage returns a user-friendly validation message
func (c *AuthController) getCustomMessage(field, tag string) string {
	messages := map[string]map[string]string{
		"Email": {
			"required": "Email is required.",
			"email":    "Please provide a valid email address.",
		},
		"Password": {
			"required": "Password is required.",
			"min":      "Password must be at least 6 characters long.",
		},
	}

	if msg, exists := messages[field][tag]; exists {
		return msg
	}
	return "Invalid input."
}
