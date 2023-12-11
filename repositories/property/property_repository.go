package property

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/repositories/property/mysql"
	"pusat-rumah-lelang-backend/requests"

	"gorm.io/gorm"
)

type IPropertyRepository interface {
	interfaces.IGetByIdGeneric[models.Property]
	interfaces.ICreateGeneric[models.Property]
	interfaces.IDeleteGeneric[models.Property]
	interfaces.IUpdateGeneric[models.Property]
	GetAll(req requests.PropertyPaginationRequest) ([]models.Property, error)
	GetBySellingStatus(id int64, page, pageSize int) ([]models.Property, error)
	GetByLocation(latitude string, longitude string, radius string) ([]models.Property, error)
	GetTotalByCategory(category string) (result []models.NameCount, err error)
}

type PropertyRepository struct {
	mysql mysql.IPropertyMysql
}

func NewPropertyRepository(mysqlDB *gorm.DB) IPropertyRepository {
	mysql := mysql.NewPropertyMysql(mysqlDB)

	return &PropertyRepository{
		mysql: mysql,
	}
}

func (r *PropertyRepository) GetTotalByCategory(category string) (result []models.NameCount, err error) {
	return r.mysql.GetTotalByCategory(category)
}

func (r *PropertyRepository) Create(data *models.Property) error {
	return r.mysql.Create(data)
}

func (r *PropertyRepository) GetAll(req requests.PropertyPaginationRequest) ([]models.Property, error) {
	return r.mysql.GetAll(req)
}

func (r *PropertyRepository) GetById(id int64) (models.Property, error) {
	return r.mysql.GetById(id)
}

func (r *PropertyRepository) Update(id int64, data *models.Property) error {
	return r.mysql.Update(id, data)
}

func (r *PropertyRepository) Delete(id int64) error {
	return r.mysql.Delete(id)
}

func (r *PropertyRepository) GetBySellingStatus(id int64, page, pageSize int) ([]models.Property, error) {
	return r.mysql.GetBySellingStatus(id, page, pageSize)
}

func (r *PropertyRepository) GetByLocation(latitude string, longitude string, radius string) ([]models.Property, error) {
	return r.mysql.GetByLocation(latitude, longitude, radius)
}
