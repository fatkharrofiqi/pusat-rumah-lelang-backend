package interfaces

import "github.com/gin-gonic/gin"

type IGetAllHandler interface {
	GetAll(c *gin.Context)
}
