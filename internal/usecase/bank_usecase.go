package usecase

import (
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/models"
	"pusat-rumah-lelang-backend/internal/repositories/bank"
)

type IBankUsecase interface {
	interfaces.IGetAllGeneric[models.Bank]
}

type BankUsecase struct {
	repo bank.IBankRepository
}

func NewBankUsecase(repo bank.IBankRepository) IBankUsecase {
	return &BankUsecase{
		repo: repo,
	}
}

func (b *BankUsecase) GetAll() ([]models.Bank, error) {
	return b.repo.GetAll()
}
