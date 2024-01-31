package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

type StandardResponse struct {
	Status int           `json:"status"`
	Error  []string      `json:"error,omitempty"`
	Data   interface{}   `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}, paging *PageMetadata) {
	c.JSON(http.StatusOK, StandardResponse{
		Status: http.StatusOK,
		Data:   data,
		Paging: paging,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message []string) {
	c.JSON(statusCode, StandardResponse{
		Status: statusCode,
		Error:  message,
		Data:   nil,
	})
}
