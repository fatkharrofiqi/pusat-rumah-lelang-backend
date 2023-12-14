package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/usecases"

	"github.com/gin-gonic/gin"
)

type IBannerHandler interface {
	interfaces.IGetAllHandler
}

type BannerHandler struct {
	usecase usecases.IBannerUsecase
}

func NewBannerHandler(usecase usecases.IBannerUsecase) IBannerHandler {
	return &BannerHandler{usecase: usecase}
}

func (h *BannerHandler) GetAll(c *gin.Context) {
	results, err := h.usecase.GetAll(0, 0)

	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(c, results, "Success")
}
