package bank

import (
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/repository/bank/mysql"

	"gorm.io/gorm"
)

type IBankRepository interface {
	interfaces.IGetAllGeneric[model.Bank]
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

func (r *BankRepository) GetAll(page, pageSize int) ([]model.Bank, error) {
	return r.mysql.GetAll(page, pageSize)
}
