package mysql

import (
	"pusat-rumah-lelang-backend/internal/models"

	"gorm.io/gorm"
)

type IPropertyMysql interface {
	Create(*models.Property) error
	Update(id int64, property *models.Property) error
	Delete(id int64) error
	GetAll() ([]models.Property, error)
	GetById(id int64) (models.Property, error)
}

type PropertyMysql struct {
	DB *gorm.DB
}

func NewPropertyMysql(db *gorm.DB) IPropertyMysql {
	return &PropertyMysql{DB: db}
}

func (r *PropertyMysql) Create(property *models.Property) error {
	return r.DB.Create(property).Error
}

func (r *PropertyMysql) Update(id int64, property *models.Property) error {
	return r.DB.Where("id = ?", id).Save(property).Error
}

func (r *PropertyMysql) Delete(id int64) error {
	data := &models.Property{}
	return r.DB.Where("id = ?", id).Delete(data).Error
}

func (r *PropertyMysql) GetAll() ([]models.Property, error) {
	var Propertys []models.Property
	if err := r.DB.Preload("SellingStatus").Preload("PhotoHouse").Preload("PhotoCertificate").Find(&Propertys).Error; err != nil {
		return nil, err
	}

	return Propertys, nil
}

func (r *PropertyMysql) GetById(id int64) (models.Property, error) {
	property := models.Property{}
	if err := r.DB.Preload("SellingStatus").Preload("PhotoHouse").Preload("PhotoCertificate").First(&property, id).Error; err != nil {
		return property, err
	}

	return property, nil
}
