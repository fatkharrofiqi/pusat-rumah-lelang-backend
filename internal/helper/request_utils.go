package helper

import (
	"github.com/gin-gonic/gin"
)

type PaginationParam struct {
	Page     int `json:"page,omitempty" form:"page"`
	PageSize int `json:"page_size,omitempty" form:"page_size"`
}

func ParsePaginationParams(c *gin.Context) (*PaginationParam, error) {
	var pagination PaginationParam

	// Parse JSON data
	if err := c.ShouldBind(&pagination); err != nil {
		return nil, err
	}

	setDefault := func(field *int, defaultValue int) {
		if *field == 0 {
			*field = defaultValue
		}
	}

	defaultPage := 1
	defaultPageSize := 10

	setDefault(&pagination.Page, defaultPage)
	setDefault(&pagination.PageSize, defaultPageSize)

	return &pagination, nil
}
