package interfaces

import "github.com/gin-gonic/gin"

type IGetByIdHandler interface {
	GetById(c *gin.Context)
}
