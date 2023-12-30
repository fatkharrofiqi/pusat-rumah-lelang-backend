package config

import "github.com/gin-gonic/gin"

func NewGin() *gin.Engine {
	gin := gin.Default()
	return gin
}
