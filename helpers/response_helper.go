package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StandardResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, StandardResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, StandardResponse{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	})
}
