package interfaces

import "github.com/gin-gonic/gin"

type ICreateHandler interface {
	Create(c *gin.Context)
}
