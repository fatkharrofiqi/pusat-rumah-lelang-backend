package handler

import (
	"fmt"
	"math"
	"net/http"
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/helper"
	"pusat-rumah-lelang-backend/internal/request"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type IBankHandler interface {
	interfaces.IGetAllHandler
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type BankHandler struct {
	Log         *logrus.Logger
	Validate    *validator.Validate
	BankUsecase usecase.IBankUsecase
}

func NewBankHandler(usecase usecase.IBankUsecase, validate *validator.Validate, log *logrus.Logger) IBankHandler {
	return &BankHandler{
		BankUsecase: usecase,
		Validate:    validate,
		Log:         log,
	}
}

func (h *BankHandler) Create(c *gin.Context) {
	request := &request.CreateBankRequest{}
	if err := c.ShouldBind(request); err != nil {
		h.Log.WithError(err).Error("failed to bind create request")
		helper.ErrorResponse(c, http.StatusBadRequest, []string{err.Error()})
		return
	}

	if err := h.Validate.Struct(request); err != nil {
		h.Log.WithError(err).Error("failed to validate request body")
		errorMessages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Validation error for %s field: %s", err.Field(), err.Tag())
			errorMessages = append(errorMessages, errorMessage)
		}
		helper.ErrorResponse(c, http.StatusBadRequest, errorMessages)
		return
	}

	if err := h.BankUsecase.Create(c, request); err != nil {
		h.Log.WithError(err).Error("failed create bank")
		helper.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	helper.SuccessResponse(c, request, nil)
}

func (h *BankHandler) Update(c *gin.Context) {
	request := &request.UpdateBankRequest{}
	if err := c.ShouldBind(request); err != nil {
		h.Log.WithError(err).Error("failed to bind update request")
		helper.ErrorResponse(c, http.StatusBadRequest, []string{err.Error()})
		return
	}

	if err := h.BankUsecase.Update(c, request); err != nil {
		h.Log.WithError(err).Error("failed create bank")
		helper.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	helper.SuccessResponse(c, request, nil)
}

func (h *BankHandler) Delete(c *gin.Context) {
	request := &request.DeleteBankRequest{}
	if err := c.ShouldBindUri(request); err != nil {
		h.Log.WithError(err).Error("failed to bind delete request")
		helper.ErrorResponse(c, http.StatusBadRequest, []string{err.Error()})
		return
	}
	fmt.Println(request)

	if err := h.BankUsecase.Delete(c, request); err != nil {
		h.Log.WithError(err).Error("failed to delete bank")
		helper.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	helper.SuccessResponse(c, request, nil)
}

func (h *BankHandler) GetAll(c *gin.Context) {
	request := &request.GetAllBankRequest{}
	if err := c.ShouldBindQuery(request); err != nil {
		h.Log.WithError(err).Error("failed bind query request")
		helper.ErrorResponse(c, http.StatusBadRequest, []string{err.Error()})
		return
	}

	result, total, err := h.BankUsecase.GetAll(c, request)
	if err != nil {
		h.Log.WithError(err).Error("failed to retrieve list bank")
		helper.ErrorResponse(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	paging := &helper.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	helper.SuccessResponse(c, result, paging)
}
