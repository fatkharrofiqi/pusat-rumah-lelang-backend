package http

import (
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Usecase struct {
	Property             usecase.IPropertyUsecase
	Bank                 usecase.IBankUsecase
	SellingStatus        usecase.ISellingStatusUsecase
	Migrate              usecase.IMigrationUsecase
	AutocompleteProperty usecase.IAutocompletePropertyUsecase
	Banner               usecase.IBannerUsecase
}

func InitUsecase(repository *Repository, db *gorm.DB, log *logrus.Logger) *Usecase {
	return &Usecase{
		Property:             usecase.NewPropertyUsecase(db, repository.Property, log),
		Bank:                 usecase.NewBankUsecase(db, repository.Bank, log),
		SellingStatus:        usecase.NewSellingStatusUsecase(db, repository.SellingStatus, log),
		Migrate:              usecase.NewMigrationUsecase(db, repository.Migration),
		AutocompleteProperty: usecase.NewAutocompleteUsecase(db, repository.AutocompleteProperty),
		Banner:               usecase.NewBannerUsecase(),
	}
}
