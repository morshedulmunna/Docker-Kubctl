package utils

import "golang.org/x/crypto/bcrypt"

// ComparePassword checks if the provided password matches the hashed password
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
