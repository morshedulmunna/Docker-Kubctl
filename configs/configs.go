package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database URL from environment variables
	dbURL := os.Getenv("DB_URL")

	// Check if DB_URL is set
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	// Connect to the database
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	// Ping the database to ensure it's reachable
	if err = DB.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	// Confirm that the connection was successful
	fmt.Println("Connected to DB")

	// Run the seed data function only if the database is not already seeded
	if !IsDatabaseSeeded() {
		Seed()
	}
}
