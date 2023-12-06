package property

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/repositories/property/mysql"

	"gorm.io/gorm"
)

type IPropertyRepository interface {
	interfaces.IGenericResource[models.Property]
	GetBySellingStatus(id int64) ([]models.Property, error)
	GetByLocation(latitude string, longitude string, radius string) ([]models.Property, error)
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

func (r *PropertyRepository) Create(data *models.Property) error {
	return r.mysql.Create(data)
}

func (r *PropertyRepository) GetAll() ([]models.Property, error) {
	return r.mysql.GetAll()
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

func (r *PropertyRepository) GetBySellingStatus(id int64) ([]models.Property, error) {
	return r.mysql.GetBySellingStatus(id)
}

func (r *PropertyRepository) GetByLocation(latitude string, longitude string, radius string) ([]models.Property, error) {
	return r.mysql.GetByLocation(latitude, longitude, radius)
}
