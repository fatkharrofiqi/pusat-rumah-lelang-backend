package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/usecases"

	"github.com/gin-gonic/gin"
)

type IBankHandler interface {
	interfaces.IGetAllHandler
}

type BankHandler struct {
	usecase usecases.IBankUsecase
}

func NewBankHandler(usecase usecases.IBankUsecase) IBankHandler {
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
