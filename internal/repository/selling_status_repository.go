package repository

import (
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/request"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ISellingStatusRepository interface {
	GetAll(db *gorm.DB, request *request.GetAllSellingStatusRequest) ([]*model.SellingStatus, error)
	Create(db *gorm.DB, request *request.CreateSellingStatusRequest) error
	Update(db *gorm.DB, id int64, request *request.UpdateSellingStatusRequest) error
	Delete(db *gorm.DB, id int64) error
	GetById(db *gorm.DB, id int64) (*model.SellingStatus, error)
	Total(db *gorm.DB) (count int64)
}

type SellingStatusRepository struct {
	Repository[model.SellingStatus]
	Log *logrus.Logger
}

func NewSellingStatusRepository(log *logrus.Logger) ISellingStatusRepository {
	return &SellingStatusRepository{Log: log}
}

func (mysql *SellingStatusRepository) Total(db *gorm.DB) (count int64) {
	db.Model(&model.SellingStatus{}).Count(&count)
	return
}

func (r *SellingStatusRepository) GetAll(db *gorm.DB, request *request.GetAllSellingStatusRequest) ([]*model.SellingStatus, error) {
	sellingStatuses := []*model.SellingStatus{}
	offset := (request.Page - 1) * request.Size
	if err := db.
		Limit(request.Size).
		Offset(offset).
		Find(&sellingStatuses).Error; err != nil {
		r.Log.WithError(err).Error("failed to get all selling status repositories")
		return nil, err
	}

	return sellingStatuses, nil
}

func (mysql *SellingStatusRepository) Create(db *gorm.DB, request *request.CreateSellingStatusRequest) error {
	return db.Create(request).Error
}

func (mysql *SellingStatusRepository) Update(db *gorm.DB, id int64, updatingStatus *request.UpdateSellingStatusRequest) error {
	return db.Where("id = ?", id).Save(updatingStatus).Error
}

func (mysql *SellingStatusRepository) Delete(db *gorm.DB, id int64) error {
	return db.Where("id = ?", id).Delete(&model.SellingStatus{}).Error
}

func (mysql *SellingStatusRepository) GetById(db *gorm.DB, id int64) (*model.SellingStatus, error) {
	data := &model.SellingStatus{}
	if err := db.Where("id = ?", id).Find(data).Error; err != nil {
		return data, err
	}

	return data, nil
}
