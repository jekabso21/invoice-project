package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the specified .env file
func LoadEnv() {
	err := godotenv.Load("../.env") // Adjust this path if necessary
	if err != nil {
		log.Println("Warning: .env file not found or failed to load")
	}
}

// GetEnv retrieves an environment variable or returns a fallback value
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
