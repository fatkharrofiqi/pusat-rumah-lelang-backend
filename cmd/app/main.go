package main

import (
	"log"
	"net/http"
	"pusat-rumah-lelang-backend/internal"
	"pusat-rumah-lelang-backend/internal/handlers"
	"pusat-rumah-lelang-backend/internal/repositories/property"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := internal.InitializeDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	// Initialize
	propertyRepository := property.NewPropertyRepository(db)
	propertyUsecase := usecase.NewPropertyUsecase(propertyRepository)
	propertyHandler := handlers.NewPropertyHandler(propertyUsecase)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "API Pusat Rumah Lelang")
	})

	// Property routes
	propertyRoutes := r.Group("/property")
	{
		propertyRoutes.POST("/", propertyHandler.Create)
		propertyRoutes.GET("/", propertyHandler.GetAll)
		propertyRoutes.GET("/:id", propertyHandler.GetById)
		propertyRoutes.PUT("/:id", propertyHandler.Update)
		propertyRoutes.DELETE("/:id", propertyHandler.Delete)
		propertyRoutes.GET("/category/:id", propertyHandler.GetByCategory)
	}

	r.Run(":8080")
}
