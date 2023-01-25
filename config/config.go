package config

import "os"

// Config is the configuration for the application
type Config struct {
	DBName string
}

// New creates a new configuration
func New() *Config {
	dbName := os.Getenv("ATOUS_DB_NAME")
	if dbName == "" {
		dbName = "atous.db"
	}
	return &Config{
		DBName: dbName,
	}
}
