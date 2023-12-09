package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/usecases"

	"github.com/gin-gonic/gin"
)

type IPropertyHandler interface {
	interfaces.IResourceHandler
	GetBySellingStatus(ctx *gin.Context)
	GetByLocation(ctx *gin.Context)
	GetTotalByCategory(ctx *gin.Context)
}

type PropertyHandler struct {
	usecase usecases.IPropertyUsecase
}

func NewPropertyHandler(usecase usecases.IPropertyUsecase) IPropertyHandler {
	return &PropertyHandler{
		usecase: usecase,
	}
}

func (h *PropertyHandler) GetTotalByCategory(ctx *gin.Context) {
	var response = make(map[string]interface{})
	category := ctx.Param("category")
	if category != "selling_status" && category != "bank" {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid category")
		return
	}

	result, err := h.usecase.GetTotalByCategory(category)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	response[category] = result
	helpers.SuccessResponse(ctx, response, "Success")
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
	pagination, err := helpers.ParsePaginationParams(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid param")
	}
	result, err := h.usecase.GetAll(pagination.Page, pagination.PageSize)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.SuccessResponse(ctx, result, "success")
}

func (h *PropertyHandler) GetById(ctx *gin.Context) {
	propertyId, err := helpers.GetID(ctx)
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
	propertyId, err := helpers.GetID(ctx)
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
	propertyId, err := helpers.GetID(ctx)
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

func (h *PropertyHandler) GetBySellingStatus(ctx *gin.Context) {
	categoryId, err := helpers.GetID(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	pagination, err := helpers.ParsePaginationParams(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid param")
	}

	result, err := h.usecase.GetBySellingStatus(int64(categoryId), pagination.Page, pagination.PageSize)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, result, "successfully retrieved")
}

func (h *PropertyHandler) GetByLocation(c *gin.Context) {
	latitude := c.DefaultQuery("latitude", "0")
	longitude := c.DefaultQuery("longitude", "0")
	radius := c.DefaultQuery("radius", "0")

	results, err := h.usecase.GetByLocation(latitude, longitude, radius)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(c, results, "successfully retrieved")
}
