package usecases

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/repositories/property"
	"pusat-rumah-lelang-backend/requests"
)

type IPropertyUsecase interface {
	interfaces.IGetByIdGeneric[models.Property]
	interfaces.ICreateGeneric[models.Property]
	interfaces.IDeleteGeneric[models.Property]
	interfaces.IUpdateGeneric[models.Property]
	GetAll(req requests.PropertyPaginationRequest) ([]models.Property, error)
	GetBySellingStatus(id int64, page, pageSize int) ([]models.Property, error)
	GetByLocation(latitude string, longitude string, radius string) ([]models.Property, error)
	GetTotalByCategory(category string) (result []models.NameCount, err error)
}

type PropertyUsecase struct {
	repo property.IPropertyRepository
}

func NewPropertyUsecase(repo property.IPropertyRepository) IPropertyUsecase {
	return &PropertyUsecase{
		repo: repo,
	}
}

func (p *PropertyUsecase) GetTotalByCategory(category string) (result []models.NameCount, err error) {
	return p.repo.GetTotalByCategory(category)
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

func (p *PropertyUsecase) GetAll(req requests.PropertyPaginationRequest) ([]models.Property, error) {
	result, err := p.repo.GetAll(req)
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

func (p *PropertyUsecase) GetBySellingStatus(id int64, page, pageSize int) ([]models.Property, error) {
	results, err := p.repo.GetBySellingStatus(id, page, pageSize)
	if err != nil {
		return results, err
	}

	return results, nil
}

func (p *PropertyUsecase) GetByLocation(latitude string, longitude string, radius string) ([]models.Property, error) {
	result, err := p.repo.GetByLocation(latitude, longitude, radius)
	if err != nil {
		return result, err
	}

	return result, nil
}
