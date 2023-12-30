package property

import (
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/repository/property/mysql"
	"pusat-rumah-lelang-backend/internal/request"

	"gorm.io/gorm"
)

type IPropertyRepository interface {
	interfaces.IGetByIdGeneric[model.Property]
	interfaces.ICreateGeneric[model.Property]
	interfaces.IDeleteGeneric[model.Property]
	interfaces.IUpdateGeneric[model.Property]
	GetAll(req request.PropertyPaginationRequest) ([]model.Property, error)
	GetBySellingStatus(id int64, page, pageSize int) ([]model.Property, error)
	GetByLocation(latitude string, longitude string, radius string) ([]model.Property, error)
	GetTotalByCategory(category string) (result []model.NameCount, err error)
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

func (r *PropertyRepository) GetTotalByCategory(category string) (result []model.NameCount, err error) {
	return r.mysql.GetTotalByCategory(category)
}

func (r *PropertyRepository) Create(data *model.Property) error {
	return r.mysql.Create(data)
}

func (r *PropertyRepository) GetAll(req request.PropertyPaginationRequest) ([]model.Property, error) {
	return r.mysql.GetAll(req)
}

func (r *PropertyRepository) GetById(id int64) (model.Property, error) {
	return r.mysql.GetById(id)
}

func (r *PropertyRepository) Update(id int64, data *model.Property) error {
	return r.mysql.Update(id, data)
}

func (r *PropertyRepository) Delete(id int64) error {
	return r.mysql.Delete(id)
}

func (r *PropertyRepository) GetBySellingStatus(id int64, page, pageSize int) ([]model.Property, error) {
	return r.mysql.GetBySellingStatus(id, page, pageSize)
}

func (r *PropertyRepository) GetByLocation(latitude string, longitude string, radius string) ([]model.Property, error) {
	return r.mysql.GetByLocation(latitude, longitude, radius)
}
