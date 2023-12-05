package routes

import (
	"pusat-rumah-lelang-backend/internal/handlers"
	"pusat-rumah-lelang-backend/internal/repositories/sellingstatus"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitSellingStatusRoute(db *gorm.DB, r *gin.Engine) {
	sellingStatusRepository := sellingstatus.NewSellingStatusRepository(db)
	sellingStatusUsecase := usecase.NewSellingStatusUsecase(sellingStatusRepository)
	sellingStatusHandler := handlers.NewSellingStatusHandler(sellingStatusUsecase)

	// Selling Status routes
	sellingStatusRoutes := r.Group("/selling_status")
	{
		sellingStatusRoutes.POST("/", sellingStatusHandler.Create)
		sellingStatusRoutes.GET("/", sellingStatusHandler.GetAll)
		sellingStatusRoutes.GET("/:id", sellingStatusHandler.GetById)
		sellingStatusRoutes.PUT("/:id", sellingStatusHandler.Update)
		sellingStatusRoutes.DELETE("/:id", sellingStatusHandler.Delete)
	}
}
