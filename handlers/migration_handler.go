package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/constants"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/usecases"

	"github.com/gin-gonic/gin"
)

type IMigrationHandler interface {
	Migrate(c *gin.Context)
}

type MigrationHandler struct {
	usecase usecases.IMigrationUsecase
}

func NewMigrationHandler(usecase usecases.IMigrationUsecase) IMigrationHandler {
	return &MigrationHandler{usecase: usecase}
}

func (h *MigrationHandler) Migrate(c *gin.Context) {
	password := c.DefaultQuery("password", "")

	if password != constants.PasswordMigration {
		helpers.ErrorResponse(c, http.StatusForbidden, "Forbidden")
		return
	}

	h.usecase.Migrate()
	helpers.SuccessResponse(c, nil, "Successfully migrate db")
}
