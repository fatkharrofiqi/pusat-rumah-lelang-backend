package interfaces

import "github.com/gin-gonic/gin"

type IUpdateHandler interface {
	Update(c *gin.Context)
}
