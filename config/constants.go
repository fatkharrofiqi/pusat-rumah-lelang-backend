package config

import (
	"os"
	"pusat-rumah-lelang-backend/constants"
	"pusat-rumah-lelang-backend/helpers"
	"strconv"
)

func LoadConstant() {
	constants.DBHost = os.Getenv("DB_HOST")
	constants.DBName = os.Getenv("DB_NAME")
	constants.DBPassword = os.Getenv("DB_PASSWORD")
	constants.DBPort = os.Getenv("DB_PORT")
	constants.DBUser = os.Getenv("DB_USER")

	constants.ServerPort = os.Getenv("SERVER_PORT")
	constants.PasswordMigration = os.Getenv("PASSWORD_MIGRATION")
	constants.MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	constants.MinioAccessKeyId = os.Getenv("MINIO_ACCESS_KEY_ID")
	constants.MinioSecretAccessKey = os.Getenv("MINIO_SECRET_ACCESS_KEY")

	// Convert the string to a boolean value
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))
	if err != nil {
		useSSL = false // or set a default value
	}
	constants.MinioUseSSL = useSSL
	constants.BucketName = os.Getenv("BUCKET_NAME")

	rootdir, err := helpers.GetRootDir()
	if err != nil {
		panic(err.Error())
	}

	constants.RootDir = rootdir
}
