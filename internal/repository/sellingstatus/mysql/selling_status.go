package mysql

import (
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/model"

	"gorm.io/gorm"
)

type ISellingStatusMysql interface {
	interfaces.IGenericResource[model.SellingStatus]
	Total() (count int64)
}

type SellingStatusMysql struct {
	db *gorm.DB
}

func NewSellingStatusMysql(db *gorm.DB) ISellingStatusMysql {
	return &SellingStatusMysql{db: db}
}

func (mysql *SellingStatusMysql) Total() (count int64) {
	mysql.db.Model(&model.SellingStatus{}).Count(&count)
	return
}

func (mysql *SellingStatusMysql) Create(sellingStatus *model.SellingStatus) error {
	return mysql.db.Create(sellingStatus).Error
}

func (mysql *SellingStatusMysql) Update(id int64, updatingStatus *model.SellingStatus) error {
	return mysql.db.Where("id = ?", id).Save(updatingStatus).Error
}

func (mysql *SellingStatusMysql) Delete(id int64) error {
	data := &model.SellingStatus{}
	return mysql.db.Where("id = ?", id).Delete(data).Error
}

func (mysql *SellingStatusMysql) GetAll(page, pageSize int) ([]model.SellingStatus, error) {
	sellingStatuses := []model.SellingStatus{}
	offset := (page - 1) * pageSize
	if err := mysql.db.
		Limit(pageSize).
		Offset(offset).
		Find(&sellingStatuses).Error; err != nil {
		return nil, err
	}

	return sellingStatuses, nil
}

func (mysql *SellingStatusMysql) GetById(id int64) (model.SellingStatus, error) {
	data := model.SellingStatus{}
	if err := mysql.db.Where("id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
