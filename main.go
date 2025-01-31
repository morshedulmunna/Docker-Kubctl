package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	config "github.com/morshedulmunna/pxomart-api/configs"
	"github.com/morshedulmunna/pxomart-api/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	config.ConnectDB()

	// Ensure the database connection is initialized
	if config.DB == nil {
		log.Fatal("Database connection is nil")
	}

	// Set up routes with the database connection
	router := routes.SetupRoutes(config.DB)

	// Start the server
	log.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
