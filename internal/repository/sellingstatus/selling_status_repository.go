package sellingstatus

import (
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/repository/sellingstatus/mysql"

	"gorm.io/gorm"
)

type ISellingStatusRepository interface {
	interfaces.IGenericResource[model.SellingStatus]
	Total() int64
}

type SellingStatusRepository struct {
	mysql mysql.ISellingStatusMysql
}

func NewSellingStatusRepository(mysqlDB *gorm.DB) ISellingStatusRepository {
	mysql := mysql.NewSellingStatusMysql(mysqlDB)

	return &SellingStatusRepository{mysql: mysql}
}

func (repo *SellingStatusRepository) Total() int64 {
	return repo.mysql.Total()
}

func (repo *SellingStatusRepository) Create(data *model.SellingStatus) error {
	return repo.mysql.Create(data)
}

func (repo *SellingStatusRepository) Update(id int64, data *model.SellingStatus) error {
	return repo.mysql.Update(id, data)
}

func (repo *SellingStatusRepository) Delete(id int64) error {
	return repo.mysql.Delete(id)
}

func (repo *SellingStatusRepository) GetAll(page, pageSize int) ([]model.SellingStatus, error) {
	return repo.mysql.GetAll(page, pageSize)
}

func (repo *SellingStatusRepository) GetById(id int64) (model.SellingStatus, error) {
	return repo.mysql.GetById(id)
}
