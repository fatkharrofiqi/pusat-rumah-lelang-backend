package config

import "pusat-rumah-lelang-backend/usecases"

type Usecase struct {
	property      usecases.IPropertyUsecase
	bank          usecases.IBankUsecase
	sellingStatus usecases.ISellingStatusUsecase
}

func InitUsecase(repo *Repository) *Usecase {
	return &Usecase{
		property:      usecases.NewPropertyUsecase(repo.property),
		bank:          usecases.NewBankUsecase(repo.bank),
		sellingStatus: usecases.NewSellingStatusUsecase(repo.sellingStatus),
	}
}
