package mysql

import (
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/model"

	"gorm.io/gorm"
)

type IBankMysql interface {
	interfaces.IGetAllGeneric[model.Bank]
}

type BankMysql struct {
	DB *gorm.DB
}

func NewBankMysql(db *gorm.DB) IBankMysql {
	return &BankMysql{DB: db}
}

func (db *BankMysql) GetAll(page, pageSize int) ([]model.Bank, error) {
	bank := []model.Bank{}
	offset := (page - 1) * pageSize
	if err := db.DB.
		Limit(pageSize).
		Offset(offset).
		Find(&bank).Error; err != nil {
		return bank, err
	}

	return bank, nil
}
