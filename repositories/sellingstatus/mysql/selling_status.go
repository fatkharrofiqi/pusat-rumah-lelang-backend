package mysql

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"

	"gorm.io/gorm"
)

type ISellingStatusMysql interface {
	interfaces.IGenericResource[models.SellingStatus]
}

type SellingStatusMysql struct {
	db *gorm.DB
}

func NewSellingStatusMysql(db *gorm.DB) ISellingStatusMysql {
	return &SellingStatusMysql{db: db}
}

func (mysql *SellingStatusMysql) Create(sellingStatus *models.SellingStatus) error {
	return mysql.db.Create(sellingStatus).Error
}

func (mysql *SellingStatusMysql) Update(id int64, updatingStatus *models.SellingStatus) error {
	return mysql.db.Where("id = ?", id).Save(updatingStatus).Error
}

func (mysql *SellingStatusMysql) Delete(id int64) error {
	data := &models.SellingStatus{}
	return mysql.db.Where("id = ?", id).Delete(data).Error
}

func (mysql *SellingStatusMysql) GetAll(page, pageSize int) ([]models.SellingStatus, error) {
	sellingStatuses := []models.SellingStatus{}
	offset := (page - 1) * pageSize
	if err := mysql.db.
		Preload("Property").
		Limit(pageSize).
		Offset(offset).
		Find(&sellingStatuses).Error; err != nil {
		return nil, err
	}

	return sellingStatuses, nil
}

func (mysql *SellingStatusMysql) GetById(id int64) (models.SellingStatus, error) {
	data := models.SellingStatus{}
	if err := mysql.db.Where("id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
