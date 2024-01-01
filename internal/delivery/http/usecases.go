package http

import (
	"pusat-rumah-lelang-backend/internal/usecase"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

func InitUsecase(repository *Repository, db *gorm.DB, log *logrus.Logger, viper *viper.Viper) *Usecase {
	return &Usecase{
		Property:             usecase.NewPropertyUsecase(db, repository.Property, log, viper),
		Bank:                 usecase.NewBankUsecase(db, repository.Bank, log),
		SellingStatus:        usecase.NewSellingStatusUsecase(db, repository.SellingStatus, log),
		Migrate:              usecase.NewMigrationUsecase(db, repository.Migration),
		AutocompleteProperty: usecase.NewAutocompleteUsecase(db, repository.AutocompleteProperty),
		Banner:               usecase.NewBannerUsecase(),
	}
}
