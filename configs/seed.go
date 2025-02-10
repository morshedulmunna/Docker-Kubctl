package config

import (
	"fmt"
	"log"

	"github.com/morshedulmunna/pxomart-api/utils"
)

// Seed function to populate the roles and users table
func Seed() {
	seedRoles()
	seedUsers()
}

// seedRoles function to insert default roles into the roles table
func seedRoles() {
	// List of default roles
	roles := []string{"admin", "user"}

	for _, role := range roles {
		// Check if the role already exists, avoid duplicates
		var count int
		err := DB.QueryRow("SELECT count(*) FROM roles WHERE name = $1", role).Scan(&count)
		if err != nil {
			log.Fatalf("Error checking role existence: %v\n", err)
		}

		if count == 0 {
			// Insert the role if it doesn't exist
			_, err := DB.Exec("INSERT INTO roles (name) VALUES ($1)", role)
			if err != nil {
				log.Fatalf("Error seeding role %s: %v\n", role, err)
			}
			fmt.Printf("Seeded role: %s\n", role)
		} else {
			fmt.Printf("Role %s already exists, skipping...\n", role)
		}
	}
}

func seedUsers() {
	// List of default users
	users := []struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
		Role      string
	}{
		{"John", "Doe", "john.doe@example.com", "password123", "admin"},
		{"Jane", "Smith", "jane.smith@example.com", "password123", "user"},
	}

	for _, user := range users {
		// Get the role ID for the given role name
		var roleID string
		err := DB.QueryRow("SELECT id FROM roles WHERE name = $1", user.Role).Scan(&roleID)
		if err != nil {
			log.Fatalf("Error fetching role ID for user %s: %v\n", user.Email, err)
		}

		// Check if the user already exists based on email to avoid duplicates
		var count int
		err = DB.QueryRow("SELECT count(*) FROM users WHERE email = $1", user.Email).Scan(&count)
		if err != nil {
			log.Fatalf("Error checking user existence: %v\n", err)
		}

		if count == 0 {
			// Hash the password before inserting it into the database
			hashedPassword := utils.HashPassword(user.Password)

			// Insert the user if it doesn't exist
			_, err := DB.Exec(`
				INSERT INTO users (role_id, first_name, last_name, email, password, status)
				VALUES ($1, $2, $3, $4, $5, 'active')`,
				roleID, user.FirstName, user.LastName, user.Email, hashedPassword)
			if err != nil {
				log.Fatalf("Error seeding user %s: %v\n", user.Email, err)
			}
			fmt.Printf("Seeded user: %s\n", user.Email)
		} else {
			fmt.Printf("User %s already exists, skipping...\n", user.Email)
		}
	}
}
