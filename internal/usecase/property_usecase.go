package usecase

import (
	"pusat-rumah-lelang-backend/internal/models"
	"pusat-rumah-lelang-backend/internal/repositories/property"
)

type IPropertyUsecase interface {
	Create(property *models.Property) error
	Update(id int64, property *models.Property) error
	Delete(id int64) error
	GetAll() ([]models.Property, error)
	GetById(id int64) (models.Property, error)
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
