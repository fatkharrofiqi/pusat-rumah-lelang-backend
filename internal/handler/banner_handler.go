package handler

import (
	"net/http"
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/helper"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type IBannerHandler interface {
	interfaces.IGetAllHandler
}

type BannerHandler struct {
	usecase usecase.IBannerUsecase
}

func NewBannerHandler(usecase usecase.IBannerUsecase) IBannerHandler {
	return &BannerHandler{usecase: usecase}
}

func (h *BannerHandler) GetAll(c *gin.Context) {
	results, err := h.usecase.GetAll(0, 0)

	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	helper.SuccessResponse(c, results, nil)
}
