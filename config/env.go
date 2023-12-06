package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	envVariable := godotenv.Load(environmentPath)
	if envVariable != nil {
		log.Fatal("Error loading .env file")
	}
}
