package bank

import (
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/models"
	"pusat-rumah-lelang-backend/internal/repositories/bank/mysql"

	"gorm.io/gorm"
)

type IBankRepository interface {
	interfaces.IGetAllGeneric[models.Bank]
}

type BankRepository struct {
	mysql mysql.IBankMysql
}

func NewBankRepository(db *gorm.DB) IBankRepository {
	mysql := mysql.NewBankMysql(db)

	return &BankRepository{
		mysql: mysql,
	}
}

func (r *BankRepository) GetAll() ([]models.Bank, error) {
	return r.mysql.GetAll()
}
