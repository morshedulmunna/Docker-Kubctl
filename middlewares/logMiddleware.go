package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs each incoming request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture the start time
		start := time.Now()

		// Log request details
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the response time and method
		log.Printf("Request %s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}
