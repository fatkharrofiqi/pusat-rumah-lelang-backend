package main

import (
	"fmt"
	"pusat-rumah-lelang-backend/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewGin()
	validate := config.NewValidator()
	log := config.NewLogger(viperConfig)
	minio := config.NewMinio(viperConfig, log)
	db := config.NewDatabase(viperConfig, log)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Storage:  minio,
		Validate: validate,
		Config:   viperConfig,
		Log:      log,
	})

	webPort := viperConfig.GetInt("SERVER_PORT")
	err := app.Run(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}
