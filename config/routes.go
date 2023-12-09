package config

import (
	"net/http"
	"pusat-rumah-lelang-backend/constants"

	"github.com/gin-gonic/gin"
)

func InitRoute(handler *Handler) {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "API Pusat Rumah Lelang")
	})

	propertyRoutes := r.Group("/property")
	{
		propertyRoutes.POST("/", handler.property.Create)
		propertyRoutes.GET("/", handler.property.GetAll)
		propertyRoutes.GET("/:id", handler.property.GetById)
		propertyRoutes.PUT("/:id", handler.property.Update)
		propertyRoutes.DELETE("/:id", handler.property.Delete)
		propertyRoutes.GET("/selling_status/:id", handler.property.GetBySellingStatus)
		propertyRoutes.GET("/location", handler.property.GetByLocation)
		propertyRoutes.GET("/total/:category", handler.property.GetTotalByCategory)
	}

	sellingStatusRoutes := r.Group("/selling_status")
	{
		sellingStatusRoutes.POST("/", handler.sellingStatus.Create)
		sellingStatusRoutes.GET("/", handler.sellingStatus.GetAll)
		sellingStatusRoutes.GET("/:id", handler.sellingStatus.GetById)
		sellingStatusRoutes.PUT("/:id", handler.sellingStatus.Update)
		sellingStatusRoutes.DELETE("/:id", handler.sellingStatus.Delete)
		sellingStatusRoutes.GET("/total", handler.sellingStatus.Total)
	}

	bankRoutes := r.Group("/bank")
	{
		bankRoutes.GET("/", handler.bank.GetAll)
	}

	migrateRoute := r.Group("/migrate")
	{
		migrateRoute.GET("/", handler.migrate.Migrate)
	}

	r.Run(":" + constants.ServerPort)
}
