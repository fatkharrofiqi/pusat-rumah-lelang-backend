package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/requests"
	"pusat-rumah-lelang-backend/usecases"

	"github.com/gin-gonic/gin"
)

type IAutocompleteHandler interface {
	GetComboProperty(c *gin.Context)
}

type AutocompleteHandler struct {
	usecase usecases.IAutocompleteUsecase
}

func NewAutocompleteHandler(usecase usecases.IAutocompleteUsecase) IAutocompleteHandler {
	return &AutocompleteHandler{usecase: usecase}
}

func (h *AutocompleteHandler) GetComboProperty(c *gin.Context) {
	pagination, err := helpers.ParsePaginationParams(c)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Invalid params pagination")
		return
	}

	request := &requests.AutocompletePaginationRequest{
		Title:    c.DefaultQuery("title", ""),
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
	}

	result, err := h.usecase.GetComboProperty(request)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to get combo property")
		return
	}

	helpers.SuccessResponse(c, result, "Successfully retrieved combo property")
}
