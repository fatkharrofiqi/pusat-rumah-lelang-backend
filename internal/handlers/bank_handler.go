package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type IBankHandler interface {
	interfaces.IGetAllHandler
}

type BankHandler struct {
	usecase usecase.IBankUsecase
}

func NewBankHandler(usecase usecase.IBankUsecase) IBankHandler {
	return &BankHandler{usecase: usecase}
}

func (h *BankHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll()
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(c, result, "Bank successfully retrieved")
}
