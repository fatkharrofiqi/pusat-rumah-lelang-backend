package handler

import (
	"net/http"
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/helper"
	"pusat-rumah-lelang-backend/internal/request"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IBankHandler interface {
	interfaces.IGetAllHandler
}

type BankHandler struct {
	Log         *logrus.Logger
	BankUsecase usecase.IBankUsecase
}

func NewBankHandler(usecase usecase.IBankUsecase, log *logrus.Logger) IBankHandler {
	return &BankHandler{
		BankUsecase: usecase,
		Log:         log,
	}
}

func (h *BankHandler) GetAll(c *gin.Context) {
	request := &request.GetAllBankRequest{}
	if err := c.ShouldBindQuery(request); err != nil {
		h.Log.WithError(err).Error("failed bind query request")
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.BankUsecase.GetAll(c, request)
	if err != nil {
		h.Log.WithError(err).Error("failed to retrieve list bank")
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(c, result, "Bank successfully retrieved")
}
