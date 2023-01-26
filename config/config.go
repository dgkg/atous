package config

import "os"

// Config is the configuration for the application
type Config struct {
	DBName       string
	GoogleAPIKey string
}

// New creates a new configuration
func New() *Config {
	dbName := os.Getenv("ATOUS_DB_NAME")
	if dbName == "" {
		dbName = "atous.db"
	}
	googleApiKey := os.Getenv("ATOUS_GOOGLE_API_KEY")
	if googleApiKey == "" {
		googleApiKey = "AIzaSyAU9ZJtU14RM2QNndxY0Z8TJ2zXLwt3Fnk"
	}
	// https://maps.googleapis.com/maps/api/geocode/json?address=Washington&key=YOUR_API_KEY

	return &Config{
		DBName:       dbName,
		GoogleAPIKey: googleApiKey,
	}
}
