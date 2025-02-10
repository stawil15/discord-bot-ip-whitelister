package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv retrieves an environment variable and panics if it's required but missing
func GetEnv(key string, required bool) string {
	value := os.Getenv(key)
	if required && value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}

func DebugMode() bool {
	value := os.Getenv("DEBUG")
	return value == "1" || value == "true"
}
