package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/usecases"

	"github.com/gin-gonic/gin"
)

type ISellingStatusHandler interface {
	interfaces.IResourceHandler
	Total(*gin.Context)
}

type SellingStatusHandler struct {
	usecase usecases.ISellingStatusUsecase
}

func NewSellingStatusHandler(usecase usecases.ISellingStatusUsecase) ISellingStatusHandler {
	return &SellingStatusHandler{usecase: usecase}
}

func (h *SellingStatusHandler) Total(ctx *gin.Context) {
	helpers.SuccessResponse(ctx, h.usecase.Total(), "Total selling status")
}

func (h *SellingStatusHandler) Create(ctx *gin.Context) {
	sellingStatus := models.SellingStatus{}
	if err := ctx.ShouldBindJSON(&sellingStatus); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.usecase.Create(&sellingStatus); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, sellingStatus, "created")
}

func (h *SellingStatusHandler) GetAll(ctx *gin.Context) {
	pagination, err := helpers.ParsePaginationParams(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid param")
	}
	result, err := h.usecase.GetAll(pagination.Page, pagination.PageSize)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch selling status")
		return
	}
	helpers.SuccessResponse(ctx, result, "success")
}

func (h *SellingStatusHandler) GetById(ctx *gin.Context) {
	sellingStatusID, err := helpers.GetID(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	result, err := h.usecase.GetById(int64(sellingStatusID))
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, result, "Success")
}

func (h *SellingStatusHandler) Update(ctx *gin.Context) {
	sellingStatusID, err := helpers.GetID(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	sellingStatus := models.SellingStatus{}
	if err := ctx.ShouldBindJSON(&sellingStatus); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.usecase.Update(int64(sellingStatusID), &sellingStatus); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, sellingStatus, "Selling status updated successfully")
}

func (h *SellingStatusHandler) Delete(ctx *gin.Context) {
	sellingStatusID, err := helpers.GetID(ctx)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.usecase.Delete(int64(sellingStatusID)); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(ctx, nil, "Selling status successfully deleted")
}
