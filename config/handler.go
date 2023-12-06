package config

import "pusat-rumah-lelang-backend/handlers"

type Handler struct {
	property      handlers.IPropertyHandler
	sellingStatus handlers.ISellingStatusHandler
	bank          handlers.IBankHandler
}

func InitHandlers(usecase *Usecase) *Handler {
	return &Handler{
		property:      handlers.NewPropertyHandler(usecase.property),
		bank:          handlers.NewBankHandler(usecase.bank),
		sellingStatus: handlers.NewSellingStatusHandler(usecase.sellingStatus),
	}
}
