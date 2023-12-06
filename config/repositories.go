package config

import (
	"pusat-rumah-lelang-backend/repositories/bank"
	"pusat-rumah-lelang-backend/repositories/property"
	"pusat-rumah-lelang-backend/repositories/sellingstatus"
)

type Repository struct {
	bank          bank.IBankRepository
	sellingStatus sellingstatus.ISellingStatusRepository
	property      property.IPropertyRepository
}

func InitRepository() *Repository {
	db := OpenDB()

	return &Repository{
		bank:          bank.NewBankRepository(db),
		property:      property.NewPropertyRepository(db),
		sellingStatus: sellingstatus.NewSellingStatusRepository(db),
	}
}
