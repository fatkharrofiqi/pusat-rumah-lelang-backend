package interfaces

import "github.com/gin-gonic/gin"

type IDeleteHandler interface {
	Delete(c *gin.Context)
}
