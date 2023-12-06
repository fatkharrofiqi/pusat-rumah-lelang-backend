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

func (db *BankMysql) GetAll() ([]models.Bank, error) {
	bank := []models.Bank{}
	if err := db.DB.Find(&bank).Error; err != nil {
		return bank, err
	}

	return bank, nil
}
