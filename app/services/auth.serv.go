package services

import (
	"errors"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/morshedulmunna/pxomart-api/app/repositories"
	"github.com/morshedulmunna/pxomart-api/utils"
)

// AuthService interface
type AuthService interface {
	Login(email, password string) (string, error)
}

// AuthServiceImpl struct
type AuthServiceImpl struct {
	authRepo repositories.AuthRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &AuthServiceImpl{authRepo: authRepo}
}

// Login authenticates the user and generates a JWT token
func (s *AuthServiceImpl) Login(email, password string) (string, error) {
	user, err := s.authRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Validate password
	if !utils.ComparePassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with a secret key
	secretKey := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
