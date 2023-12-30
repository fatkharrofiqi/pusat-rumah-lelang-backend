package http

import (
	"pusat-rumah-lelang-backend/internal/handler"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Property      handler.IPropertyHandler
	SellingStatus handler.ISellingStatusHandler
	Bank          handler.IBankHandler
	Migrate       handler.IMigrationHandler
	Autocomplete  handler.IAutocompleteHandler
	Banner        handler.IBannerHandler
}

func InitHandler(usecase *Usecase, log *logrus.Logger) *Handler {
	return &Handler{
		Property:      handler.NewPropertyHandler(usecase.Property, log),
		Bank:          handler.NewBankHandler(usecase.Bank, log),
		SellingStatus: handler.NewSellingStatusHandler(usecase.SellingStatus),
		Migrate:       handler.NewMigrationHandler(usecase.Migrate),
		Autocomplete:  handler.NewAutocompleteHandler(usecase.AutocompleteProperty),
		Banner:        handler.NewBannerHandler(usecase.Banner),
	}
}
