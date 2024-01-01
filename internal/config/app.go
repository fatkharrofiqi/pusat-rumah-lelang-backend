package config

import (
	"pusat-rumah-lelang-backend/internal/delivery/http"
	"pusat-rumah-lelang-backend/internal/delivery/http/route"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *gin.Engine
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Storage  *minio.Client
}

func Bootstrap(config *BootstrapConfig) {
	repository := http.InitRepository(config.DB, config.Log)
	usecase := http.InitUsecase(repository, config.DB, config.Log, config.Config)
	handler := http.InitHandler(usecase, config.Log)

	routeConfig := &route.RouteConfig{
		App:     config.App,
		Handler: handler,
	}

	routeConfig.Setup()
}
