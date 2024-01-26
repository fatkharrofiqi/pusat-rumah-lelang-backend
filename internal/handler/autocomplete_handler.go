package handler

import (
	"net/http"
	"pusat-rumah-lelang-backend/internal/helper"
	"pusat-rumah-lelang-backend/internal/request"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type IAutocompleteHandler interface {
	GetComboProperty(c *gin.Context)
}

type AutocompleteHandler struct {
	AutoCompletePropertyUsecase usecase.IAutocompletePropertyUsecase
}

func NewAutocompleteHandler(autocompletePropertyUsecase usecase.IAutocompletePropertyUsecase) IAutocompleteHandler {
	return &AutocompleteHandler{AutoCompletePropertyUsecase: autocompletePropertyUsecase}
}

func (h *AutocompleteHandler) GetComboProperty(c *gin.Context) {
	request := &request.GetComboPropertyRequest{}
	if err := c.ShouldBind(request); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "invalid params pagination")
		return
	}

	result, err := h.AutoCompletePropertyUsecase.GetComboProperty(c, request)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "failed to get combo property")
		return
	}

	helper.SuccessResponse(c, result, nil)
}
