package http

import (
	"pusat-rumah-lelang-backend/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	Bank                 repository.IBankRepository
	SellingStatus        repository.ISellingStatusRepository
	Property             repository.IPropertyRepository
	Migration            repository.IMigrationRepository
	AutocompleteProperty repository.IAutocompletePropertyRepository
}

func InitRepository(db *gorm.DB, log *logrus.Logger) *Repository {
	return &Repository{
		Bank:                 repository.NewBankRepository(log),
		Property:             repository.NewPropertyRepository(log),
		SellingStatus:        repository.NewSellingStatusRepository(log),
		Migration:            repository.NewMigrationRepository(log),
		AutocompleteProperty: repository.NewAutocompletePropertyRepository(log),
	}
}
