/*package config

import (
	"github.com/joho/godotenv"
	"log"
)

//This function loads environmental variables from .env file

func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

}*/



package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv tries to load .env for local development,
// but won't crash if the file doesn't exist (like on Render).
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables.")
	}
}

// GetEnv reads an environment variable or returns a fallback if not set.
func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
