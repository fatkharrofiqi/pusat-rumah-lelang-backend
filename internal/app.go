package internal

import (
	"log"
	"net/http"
	"pusat-rumah-lelang-backend/config"
	"pusat-rumah-lelang-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func App() {
	db, err := config.InitializeDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "API Pusat Rumah Lelang")
	})

	routes.InitPropertyRoute(db, r)
	routes.InitSellingStatusRoute(db, r)

	r.Run(":8080")
}
