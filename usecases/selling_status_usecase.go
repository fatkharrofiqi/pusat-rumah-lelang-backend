package usecases

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/repositories/sellingstatus"
)

type ISellingStatusUsecase interface {
	interfaces.IGenericResource[models.SellingStatus]
}

type SellingStatusUsecase struct {
	repo sellingstatus.ISellingStatusRepository
}

func NewSellingStatusUsecase(repo sellingstatus.ISellingStatusRepository) ISellingStatusUsecase {
	return &SellingStatusUsecase{repo: repo}
}

func (u *SellingStatusUsecase) Create(sellingStatus *models.SellingStatus) error {
	return u.repo.Create(sellingStatus)
}

func (u *SellingStatusUsecase) Update(id int64, updatingStatus *models.SellingStatus) error {
	return u.repo.Update(id, updatingStatus)
}

func (u *SellingStatusUsecase) Delete(id int64) error {
	return u.repo.Delete(id)
}

func (u *SellingStatusUsecase) GetAll() ([]models.SellingStatus, error) {
	results, err := u.repo.GetAll()
	if err != nil {
		return results, err
	}

	return results, nil
}

func (u *SellingStatusUsecase) GetById(id int64) (models.SellingStatus, error) {
	result, err := u.repo.GetById(id)
	if err != nil {
		return result, err
	}

	return result, nil
}
