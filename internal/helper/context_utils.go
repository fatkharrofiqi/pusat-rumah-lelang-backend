package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetID(ctx *gin.Context) (int64, error) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}
