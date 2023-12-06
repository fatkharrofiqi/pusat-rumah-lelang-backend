package routes

import (
	"pusat-rumah-lelang-backend/internal/handlers"
	"pusat-rumah-lelang-backend/internal/repositories/property"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPropertyRoute(db *gorm.DB, r *gin.Engine) {
	// Initialize
	propertyRepository := property.NewPropertyRepository(db)
	propertyUsecase := usecase.NewPropertyUsecase(propertyRepository)
	propertyHandler := handlers.NewPropertyHandler(propertyUsecase)

	// Property routes
	propertyRoutes := r.Group("/property")
	{
		propertyRoutes.POST("/", propertyHandler.Create)
		propertyRoutes.GET("/", propertyHandler.GetAll)
		propertyRoutes.GET("/:id", propertyHandler.GetById)
		propertyRoutes.PUT("/:id", propertyHandler.Update)
		propertyRoutes.DELETE("/:id", propertyHandler.Delete)
		propertyRoutes.GET("/selling_status/:id", propertyHandler.GetBySellingStatus)
		propertyRoutes.GET("/location", propertyHandler.GetByLocation)
	}
}
