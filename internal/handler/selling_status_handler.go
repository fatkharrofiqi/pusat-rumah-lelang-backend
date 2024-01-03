package handler

import (
	"net/http"
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/helper"
	"pusat-rumah-lelang-backend/internal/request"
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ISellingStatusHandler interface {
	interfaces.IResourceHandler
	Total(*gin.Context)
}

type SellingStatusHandler struct {
	SellingStatusUsecase usecase.ISellingStatusUsecase
}

func NewSellingStatusHandler(sellingStatusUsecase usecase.ISellingStatusUsecase) ISellingStatusHandler {
	return &SellingStatusHandler{SellingStatusUsecase: sellingStatusUsecase}
}

func (h *SellingStatusHandler) Total(ctx *gin.Context) {
	helper.SuccessResponse(ctx, h.SellingStatusUsecase.Total(ctx), "Total selling status")
}

func (h *SellingStatusHandler) Create(ctx *gin.Context) {
	createSellingStatus := &request.CreateSellingStatusRequest{}
	if err := ctx.ShouldBindJSON(createSellingStatus); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.SellingStatusUsecase.Create(ctx, createSellingStatus); err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, createSellingStatus, "created")
}

func (h *SellingStatusHandler) GetAll(ctx *gin.Context) {
	request := &request.GetAllSellingStatusRequest{}
	if err := ctx.ShouldBind(request); err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid param")
		return
	}

	result, err := h.SellingStatusUsecase.GetAll(ctx, request)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch selling status")
		return
	}

	helper.SuccessResponse(ctx, result, "success")
}

func (h *SellingStatusHandler) GetById(ctx *gin.Context) {
	sellingStatusID, err := helper.GetID(ctx)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	result, err := h.SellingStatusUsecase.GetById(ctx, int64(sellingStatusID))
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, result, "Success")
}

func (h *SellingStatusHandler) Update(ctx *gin.Context) {
	sellingStatusID, err := helper.GetID(ctx)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	request := &request.UpdateSellingStatusRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.SellingStatusUsecase.Update(ctx, int64(sellingStatusID), request); err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, request, "Selling status updated successfully")
}

func (h *SellingStatusHandler) Delete(ctx *gin.Context) {
	sellingStatusID, err := helper.GetID(ctx)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.SellingStatusUsecase.Delete(ctx, int64(sellingStatusID)); err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(ctx, nil, "Selling status successfully deleted")
}
