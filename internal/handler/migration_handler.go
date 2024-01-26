package handler

import (
	"net/http"
	"pusat-rumah-lelang-backend/internal/helper"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type IMigrationHandler interface {
	Migrate(c *gin.Context)
}

type MigrationHandler struct {
	MigrationUseecase usecase.IMigrationUsecase
}

func NewMigrationHandler(usecase usecase.IMigrationUsecase) IMigrationHandler {
	return &MigrationHandler{MigrationUseecase: usecase}
}

func (h *MigrationHandler) Migrate(c *gin.Context) {
	password := c.DefaultQuery("password", "")

	if password != "" {
		helper.ErrorResponse(c, http.StatusForbidden, "Forbidden")
		return
	}

	h.MigrationUseecase.Migrate(c)
	helper.SuccessResponse(c, nil, nil)
}
