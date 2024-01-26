package repository

import (
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/request"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IBankRepository interface {
	GetAll(db *gorm.DB, request *request.GetAllBankRequest) ([]*model.Bank, int64, error)
}

type BankRepository struct {
	Repository[model.Bank]
	Log *logrus.Logger
}

func NewBankRepository(log *logrus.Logger) IBankRepository {
	return &BankRepository{
		Log: log,
	}
}

func (r *BankRepository) GetAll(db *gorm.DB, request *request.GetAllBankRequest) ([]*model.Bank, int64, error) {
	result := []*model.Bank{}
	offset := (request.Page - 1) * request.Size
	if err := db.
		Limit(request.Size).
		Offset(offset).
		Find(&result).Error; err != nil {
		r.Log.WithError(err).Error("failed to get all bank repositories")
		return result, 0, err
	}

	var total int64 = 0
	if err := db.Model(&model.Bank{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return result, total, nil
}
