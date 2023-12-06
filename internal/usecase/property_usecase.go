package usecase

import (
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/models"
	"pusat-rumah-lelang-backend/internal/repositories/property"
)

type IPropertyUsecase interface {
	interfaces.IGenericResource[models.Property]
	GetBySellingStatus(id int64) (models.Property, error)
	GetByLocation(latitude string, longitude string) ([]models.Property, error)
}

type PropertyUsecase struct {
	repo property.IPropertyRepository
}

func NewPropertyUsecase(repo property.IPropertyRepository) IPropertyUsecase {
	return &PropertyUsecase{
		repo: repo,
	}
}

func (p *PropertyUsecase) Create(property *models.Property) error {
	if err := p.repo.Create(property); err != nil {
		return err
	}

	return nil
}

func (p *PropertyUsecase) Update(id int64, property *models.Property) error {
	if err := p.repo.Update(id, property); err != nil {
		return err
	}

	return nil
}

func (p *PropertyUsecase) Delete(id int64) error {
	if err := p.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (p *PropertyUsecase) GetAll() ([]models.Property, error) {
	result, err := p.repo.GetAll()
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *PropertyUsecase) GetById(id int64) (models.Property, error) {
	result, err := p.repo.GetById(id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *PropertyUsecase) GetBySellingStatus(id int64) (models.Property, error) {
	result, err := p.repo.GetBySellingStatus(id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *PropertyUsecase) GetByLocation(latitude string, longitude string) ([]models.Property, error) {
	result, err := p.repo.GetByLocation(latitude, longitude)
	if err != nil {
		return result, err
	}

	return result, nil
}
