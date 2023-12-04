package main

import (
	"log"
	"net/http"
	"pusat-rumah-lelang-backend/internal"
	"pusat-rumah-lelang-backend/internal/handlers"
	"pusat-rumah-lelang-backend/internal/repositories/property"
	"pusat-rumah-lelang-backend/internal/usecase"
	"pusat-rumah-lelang-backend/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := internal.InitializeDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	// Run the migration
	if err := migrations.RunMigrations(db); err != nil {
		panic("Failed to migrate table")
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
