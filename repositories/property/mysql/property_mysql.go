package mysql

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"

	"gorm.io/gorm"
)

type IPropertyMysql interface {
	interfaces.IGenericResource[models.Property]
	GetBySellingStatus(id int64) (property []models.Property, err error)
	GetByLocation(latitude string, longitude string, radius string) ([]models.Property, error)
}

type PropertyMysql struct {
	db *gorm.DB
}

func NewPropertyMysql(db *gorm.DB) IPropertyMysql {
	return &PropertyMysql{db: db}
}

func (r *PropertyMysql) Create(property *models.Property) error {
	return r.db.Create(property).Error
}

func (r *PropertyMysql) Update(id int64, property *models.Property) error {
	return r.db.Where("id = ?", id).Save(property).Error
}

func (r *PropertyMysql) Delete(id int64) error {
	data := &models.Property{}
	return r.db.Where("id = ?", id).Delete(data).Error
}

func (r *PropertyMysql) GetAll() ([]models.Property, error) {
	var Propertys []models.Property
	if err := r.db.Preload("SellingStatus").Preload("PhotoHouse").Preload("PhotoCertificate").Preload("RoadAccess").Find(&Propertys).Error; err != nil {
		return nil, err
	}

	return Propertys, nil
}

func (r *PropertyMysql) GetById(id int64) (models.Property, error) {
	property := models.Property{}
	if err := r.db.Preload("SellingStatus").Preload("PhotoHouse").Preload("PhotoCertificate").Preload("RoadAccess").First(&property, id).Error; err != nil {
		return property, err
	}

	return property, nil
}

func (r *PropertyMysql) GetBySellingStatus(id int64) (property []models.Property, err error) {
	if err := r.db.Where("selling_status_id = ?", id).Preload("SellingStatus").Preload("PhotoHouse").Preload("PhotoCertificate").Preload("RoadAccess").Find(&property, id).Error; err != nil {
		return property, err
	}

	return property, nil
}

func (r *PropertyMysql) GetByLocation(latitude string, longitude string, radius string) ([]models.Property, error) {
	properties := []models.Property{}
	sql := `SELECT
						*,
						(
							6371 *
							acos(
								cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) +
									sin(radians(?)) * sin(radians(latitude))
							)
						) as distance
					FROM
						properties p
					HAVING
						distance <= ?
					ORDER BY
						distance asc
					`
	if err := r.db.Raw(sql, latitude, longitude, latitude, radius).
		Preload("Bank").
		Preload("SellingStatus").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Find(&properties).Error; err != nil {
		return properties, err
	}

	return properties, nil
}
