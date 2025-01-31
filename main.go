package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/morshedulmunna/gogenz-library"
	config "github.com/morshedulmunna/pxomart-api/configs"
	"github.com/morshedulmunna/pxomart-api/routes"
)

var logger *log.Logger

func init() {
	// Set up logging using the utility function
	logger = gogenz.SetupLogging()

}

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

	// Set up routes
	router := routes.SetupRoutes(config.DB)

	// Wrap the router with the rate-limiting middleware
	limitedRouter := gogenz.RateLimitMiddleware(router)

	// Start the server
	logger.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", limitedRouter)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
