package request

import "pusat-rumah-lelang-backend/internal/model"

type PropertyResponse struct {
	model.Property
}

type CreatePropertyRequest struct {
	model.Property
}

type UpdatePropertyRequest struct {
	UserId    string `json:"-" validate:"required"`
	ID        string `json:"-" validate:"required,max=100,uuid"`
	FirstName string `json:"first_name" validate:"required,max=100"`
	LastName  string `json:"last_name" validate:"max=100"`
	Email     string `json:"email" validate:"max=200,email"`
	Phone     string `json:"phone" validate:"max=20"`
}

type GetAllPropertyRequest struct {
	Query string `form:"query"`
	Page  int    `form:"page,default=1"`
	Size  int    `form:"size,default=10"`
}

type GetBySellingStatusRequest struct {
	Page int `form:"page,default=1"`
	Size int `form:"size,default=10"`
}

type GetByLocationRequest struct {
	Latitude  string `form:"latitude" validate:"required"`
	Longitude string `form:"longitude" validate:"required"`
	Radius    string `form:"radius" validate:"required"`
}

type SearchPropertyRequest struct {
	UserId string `json:"-" validate:"required"`
	Name   string `json:"name" validate:"max=100"`
	Email  string `json:"email" validate:"max=200"`
	Phone  string `json:"phone" validate:"max=20"`
	Page   int    `json:"page" validate:"min=1"`
	Size   int    `json:"size" validate:"min=1,max=100"`
}

type GetPropertyRequest struct {
	UserId string `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}

type DeletePropertyRequest struct {
	UserId string `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}

type PropertyPaginationRequest struct {
	Query    string `json:"query"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
