package config

import (
	"pusat-rumah-lelang-backend/repositories/autocomplete"
	"pusat-rumah-lelang-backend/repositories/bank"
	"pusat-rumah-lelang-backend/repositories/migration"
	"pusat-rumah-lelang-backend/repositories/property"
	"pusat-rumah-lelang-backend/repositories/sellingstatus"
)

type Repository struct {
	bank          bank.IBankRepository
	sellingStatus sellingstatus.ISellingStatusRepository
	property      property.IPropertyRepository
	migrate       migration.IMigrationRepository
	autocomplete  autocomplete.IAutocompleteRepository
}

func InitRepository() *Repository {
	db := OpenDB()

	return &Repository{
		bank:          bank.NewBankRepository(db),
		property:      property.NewPropertyRepository(db),
		sellingStatus: sellingstatus.NewSellingStatusRepository(db),
		migrate:       migration.NewMigrationRepository(db),
		autocomplete:  autocomplete.NewAutocompleteRepository(db),
	}
}
