package request

import "pusat-rumah-lelang-backend/internal/model"

type SellingStatusResponse struct {
	model.SellingStatus
}

type CreateSellingStatusRequest struct {
	model.SellingStatus
}

type UpdateSellingStatusRequest struct {
	UserId    string `json:"-" validate:"required"`
	ID        string `json:"-" validate:"required,max=100,uuid"`
	FirstName string `json:"first_name" validate:"required,max=100"`
	LastName  string `json:"last_name" validate:"max=100"`
	Email     string `json:"email" validate:"max=200,email"`
	Phone     string `json:"phone" validate:"max=20"`
}

type GetAllSellingStatusRequest struct {
	Page int `form:"page,default=1"`
	Size int `form:"size,default=10"`
}

type SearchSellingStatusRequest struct {
	UserId string `json:"-" validate:"required"`
	Name   string `json:"name" validate:"max=100"`
	Email  string `json:"email" validate:"max=200"`
	Phone  string `json:"phone" validate:"max=20"`
	Page   int    `json:"page" validate:"min=1"`
	Size   int    `json:"size" validate:"min=1,max=100"`
}

type GetSellingStatusRequest struct {
	UserId string `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteSellingStatusRequest struct {
	UserId string `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}
