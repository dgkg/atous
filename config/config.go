package config

import "os"

// Config is the configuration for the application
type Config struct {
	DBName       string
	GoogleAPIKey string
	JWTKeySign   []byte
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

	JWTkey := os.Getenv("ATOUS_JWT_KEY")
	if JWTkey == "" {
		JWTkey = "test"
	}
	return &Config{
		DBName:       dbName,
		GoogleAPIKey: googleApiKey,
		JWTKeySign:   []byte(JWTkey),
	}
}
