package handlers

import (
	"net/http"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/models"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ISellingStatusHandler interface {
	interfaces.IResourceHandler
}

type SellingStatusHandler struct {
	usecase usecase.ISellingStatusUsecase
}

func NewSellingStatusHandler(usecase usecase.ISellingStatusUsecase) ISellingStatusHandler {
	return &SellingStatusHandler{usecase: usecase}
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
	result, err := h.usecase.GetAll()
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
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
