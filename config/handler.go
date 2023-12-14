package config

import "pusat-rumah-lelang-backend/handlers"

type Handler struct {
	property      handlers.IPropertyHandler
	sellingStatus handlers.ISellingStatusHandler
	bank          handlers.IBankHandler
	migrate       handlers.IMigrationHandler
	autocomplete  handlers.IAutocompleteHandler
	banner        handlers.IBannerHandler
}

func InitHandlers(usecase *Usecase) *Handler {
	return &Handler{
		property:      handlers.NewPropertyHandler(usecase.property),
		bank:          handlers.NewBankHandler(usecase.bank),
		sellingStatus: handlers.NewSellingStatusHandler(usecase.sellingStatus),
		migrate:       handlers.NewMigrationHandler(usecase.migrate),
		autocomplete:  handlers.NewAutocompleteHandler(usecase.autocomplete),
		banner:        handlers.NewBannerHandler(usecase.banner),
	}
}
