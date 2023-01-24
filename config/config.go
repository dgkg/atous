package config

import "os"

// Config is the configuration for the application
type Config struct {
	DBName string
}

// New creates a new configuration
func New() *Config {
	return &Config{
		DBName: os.Getenv("ATOUS_DB_NAME"),
	}
}
