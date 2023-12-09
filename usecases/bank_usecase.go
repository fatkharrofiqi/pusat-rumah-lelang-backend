package usecases

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/repositories/bank"
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

func (b *BankUsecase) GetAll(page, pageSize int) ([]models.Bank, error) {
	return b.repo.GetAll(page, pageSize)
}
