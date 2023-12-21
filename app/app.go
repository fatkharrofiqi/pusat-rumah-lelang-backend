package app

import (
	"log"
	"pusat-rumah-lelang-backend/config"
)

func StartApp() {
	config.LoadEnv()
	config.LoadConstant()
	minio, err := config.NewMinioStorage()
	if err != nil {
		log.Fatalln(err)
	}

	repository := config.InitRepository()
	usecase := config.InitUsecase(repository, minio)
	handler := config.InitHandlers(usecase)

	config.InitRoute(handler)
}
