package mysql

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"

	"gorm.io/gorm"
)

type IBankMysql interface {
	interfaces.IGetAllGeneric[models.Bank]
}

type BankMysql struct {
	DB *gorm.DB
}

func NewBankMysql(db *gorm.DB) IBankMysql {
	return &BankMysql{DB: db}
}

func (db *BankMysql) GetAll(page, pageSize int) ([]models.Bank, error) {
	bank := []models.Bank{}
	offset := (page - 1) * pageSize
	if err := db.DB.
		Limit(pageSize).
		Offset(offset).
		Find(&bank).Error; err != nil {
		return bank, err
	}

	return bank, nil
}
