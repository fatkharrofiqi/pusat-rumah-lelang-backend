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

type IPropertyHandler interface {
	interfaces.IResourceHandler
	GetBySellingStatus(ctx *gin.Context)
	GetByLocation(ctx *gin.Context)
	GetTotalByCategory(ctx *gin.Context)
}

type PropertyHandler struct {
	PropertyUsecase usecase.IPropertyUsecase
	Log             *logrus.Logger
}

func NewPropertyHandler(usecase usecase.IPropertyUsecase, log *logrus.Logger) IPropertyHandler {
	return &PropertyHandler{
		PropertyUsecase: usecase,
		Log:             log,
	}
}

func (h *PropertyHandler) GetTotalByCategory(ctx *gin.Context) {
	var response = make(map[string]interface{})
	category := ctx.Param("category")
	if category != "selling_status" && category != "bank" {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid category")
		return
	}

	result, err := h.PropertyUsecase.GetTotalByCategory(ctx, category)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	response[category] = result
	helper.SuccessResponse(ctx, response, "Success")
}

func (h *PropertyHandler) Create(ctx *gin.Context) {
	request := &request.CreatePropertyRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.PropertyUsecase.Create(ctx, request); err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, request, "created")
}

func (h *PropertyHandler) GetAll(ctx *gin.Context) {
	request := &request.GetAllPropertyRequest{}
	if err := ctx.ShouldBind(request); err != nil {
		h.Log.WithError(err).Error("failed to parse request get all properties")
		helper.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.PropertyUsecase.GetAll(ctx, request)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	helper.SuccessResponse(ctx, result, "success")
}

func (h *PropertyHandler) GetById(ctx *gin.Context) {
	propertyId, err := helper.GetID(ctx)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	result, err := h.PropertyUsecase.GetById(ctx, int64(propertyId))
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, result, "Success")
}

func (h *PropertyHandler) Update(ctx *gin.Context) {
	propertyId, err := helper.GetID(ctx)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	property := &request.UpdatePropertyRequest{}
	if err := ctx.ShouldBindJSON(property); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.PropertyUsecase.Update(ctx, int64(propertyId), property); err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, property, "Property updated successfully")
}

func (h *PropertyHandler) Delete(ctx *gin.Context) {
	propertyId, err := helper.GetID(ctx)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.PropertyUsecase.Delete(ctx, int64(propertyId)); err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, nil, "Property successfully deleted")
}

func (h *PropertyHandler) GetBySellingStatus(ctx *gin.Context) {
	categoryId, err := helper.GetID(ctx)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	request := &request.GetBySellingStatusRequest{}
	if err := ctx.ShouldBind(request); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.PropertyUsecase.GetBySellingStatus(ctx, int64(categoryId), request)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, result, "successfully retrieved")
}

func (h *PropertyHandler) GetByLocation(c *gin.Context) {
	request := &request.GetByLocationRequest{}
	if err := c.ShouldBind(request); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	results, err := h.PropertyUsecase.GetByLocation(c, request)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(c, results, "successfully retrieved")
}
