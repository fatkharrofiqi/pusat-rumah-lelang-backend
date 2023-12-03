package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/internal/models"
	"pusat-rumah-lelang-backend/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IPropertyHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetByCategory(ctx *gin.Context)
}

type PropertyHandler struct {
	usecase usecase.IPropertyUsecase
}

func NewPropertyHandler(usecase usecase.IPropertyUsecase) IPropertyHandler {
	return &PropertyHandler{
		usecase: usecase,
	}
}

func getPropertyID(ctx *gin.Context) (int64, error) {
	propertyIdStr := ctx.Param("id")
	propertyId, err := strconv.Atoi(propertyIdStr)
	if err != nil {
		return 0, err
	}
	return int64(propertyId), nil
}

func (h *PropertyHandler) Create(ctx *gin.Context) {
	property := models.Property{}
	if err := ctx.ShouldBindJSON(&property); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.usecase.Create(&property); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, property, "created")
}

func (h *PropertyHandler) GetAll(ctx *gin.Context) {
	result, err := h.usecase.GetAll()
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.SuccessResponse(ctx, result, "success")
}

func (h *PropertyHandler) GetById(ctx *gin.Context) {
	propertyId, err := getPropertyID(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	result, err := h.usecase.GetById(int64(propertyId))
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, result, "Success")
}

func (h *PropertyHandler) Update(ctx *gin.Context) {
	propertyId, err := getPropertyID(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	property := models.Property{}
	if err := ctx.ShouldBindJSON(&property); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.usecase.Update(int64(propertyId), &property); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, property, "Property updated successfully")
}

func (h *PropertyHandler) Delete(ctx *gin.Context) {
	propertyId, err := getPropertyID(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.usecase.Delete(int64(propertyId)); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, nil, "Property successfully deleted")
}

func (h *PropertyHandler) GetByCategory(ctx *gin.Context) {

}
