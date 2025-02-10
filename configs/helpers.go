package config

import "log"

// Function to check if the database has been seeded
func IsDatabaseSeeded() bool {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM roles").Scan(&count)
	if err != nil {
		log.Fatalf("Error checking if database is seeded: %v\n", err)
	}

	// If count is greater than 0, assume that the database is seeded
	return count > 0
}
