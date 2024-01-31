package request

import "pusat-rumah-lelang-backend/internal/model"

type BankResponse struct {
	model.Bank
}

type CreateBankRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateBankRequest struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type GetAllBankRequest struct {
	SearchBankRequest
	Page int `form:"page,default=1"`
	Size int `form:"size,default=10"`
}

type SearchBankRequest struct {
	Name string `form:"name"`
}

type GetBankRequest struct {
	UserId string `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteBankRequest struct {
	ID int `uri:"id" validate:"required"`
}
