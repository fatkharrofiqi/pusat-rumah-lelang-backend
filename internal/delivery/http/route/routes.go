package route

import (
	"pusat-rumah-lelang-backend/internal/delivery/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App     *gin.Engine
	Handler *http.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupMiddleware()
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupMiddleware() {
	c.App.Use(cors.Default())
}

func (c *RouteConfig) SetupGuestRoute() {

}

func (c *RouteConfig) SetupAuthRoute() {
	propertyRoutes := c.App.Group("/property")
	{
		propertyRoutes.POST("/", c.Handler.Property.Create)
		propertyRoutes.GET("/", c.Handler.Property.GetAll)
		propertyRoutes.GET("/:id", c.Handler.Property.GetById)
		propertyRoutes.PUT("/:id", c.Handler.Property.Update)
		propertyRoutes.DELETE("/:id", c.Handler.Property.Delete)
		propertyRoutes.GET("/selling_status/:id", c.Handler.Property.GetBySellingStatus)
		propertyRoutes.GET("/location", c.Handler.Property.GetByLocation)
		propertyRoutes.GET("/total/:category", c.Handler.Property.GetTotalByCategory)
	}

	sellingStatusRoutes := c.App.Group("/selling_status")
	{
		sellingStatusRoutes.POST("/", c.Handler.SellingStatus.Create)
		sellingStatusRoutes.GET("/", c.Handler.SellingStatus.GetAll)
		sellingStatusRoutes.GET("/:id", c.Handler.SellingStatus.GetById)
		sellingStatusRoutes.PUT("/:id", c.Handler.SellingStatus.Update)
		sellingStatusRoutes.DELETE("/:id", c.Handler.SellingStatus.Delete)
		sellingStatusRoutes.GET("/total", c.Handler.SellingStatus.Total)
	}

	bankRoutes := c.App.Group("/bank")
	{
		bankRoutes.GET("/", c.Handler.Bank.GetAll)
	}

	migrateRoute := c.App.Group("/migrate")
	{
		migrateRoute.GET("/", c.Handler.Migrate.Migrate)
	}

	autoCompleteRoute := c.App.Group("/autocomplete")
	{
		autoCompleteRoute.GET("/property", c.Handler.Autocomplete.GetComboProperty)
	}

	bannerRoute := c.App.Group("/banner")
	{
		bannerRoute.GET("/", c.Handler.Banner.GetAll)
	}
}
