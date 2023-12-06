package app

import "pusat-rumah-lelang-backend/config"

func StartApp() {
	config.LoadEnv()
	config.LoadConstant()

	repository := config.InitRepository()
	usecase := config.InitUsecase(repository)
	handler := config.InitHandlers(usecase)

	config.InitRoute(handler)
}
