package config

import (
	"log"
	"path/filepath"
	"pusat-rumah-lelang-backend/helpers"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	rootdir, err := helpers.GetRootDir()
	if err != nil {
		panic(err.Error())
	}
	environmentPath := filepath.Join(rootdir, ".env")
	envVariable := godotenv.Load(environmentPath)
	if envVariable != nil {
		log.Fatal("Error loading .env file")
	}
}
