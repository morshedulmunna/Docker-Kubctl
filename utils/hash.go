package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ComparePassword checks if the provided password matches the hashed password
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// HashPassword hashes the provided password using bcrypt
func HashPassword(password string) string {
	// Generate a hashed version of the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	return string(hashedPassword)
}
