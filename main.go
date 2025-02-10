package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/morshedulmunna/gogenz-library"
	config "github.com/morshedulmunna/pxomart-api/configs"
	"github.com/morshedulmunna/pxomart-api/routes"
)

func init() {
	gogenz.SetupLogging()
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
	router := routes.SetupRoutes()

	// Wrap the router with the rate-limiting middleware
	limitedRouter := gogenz.RateLimitMiddleware(router)

	// Start the server
	fmt.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", limitedRouter)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
