package config

import (
	"os"
	"pusat-rumah-lelang-backend/constants"
)

func LoadConstant() {
	constants.DBHost = os.Getenv("DB_HOST")
	constants.DBName = os.Getenv("DB_NAME")
	constants.DBPassword = os.Getenv("DB_PASSWORD")
	constants.DBPort = os.Getenv("DB_PORT")
	constants.DBUser = os.Getenv("DB_USER")

	constants.ServerPort = os.Getenv("SERVER_PORT")
}
