package sellingstatus

import (
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/repositories/sellingstatus/mysql"

	"gorm.io/gorm"
)

type ISellingStatusRepository interface {
	interfaces.IGenericResource[models.SellingStatus]
}

type SellingStatusRepository struct {
	mysql mysql.ISellingStatusMysql
}

func NewSellingStatusRepository(mysqlDB *gorm.DB) ISellingStatusRepository {
	mysql := mysql.NewSellingStatusMysql(mysqlDB)

	return &SellingStatusRepository{mysql: mysql}
}

func (repo *SellingStatusRepository) Create(data *models.SellingStatus) error {
	return repo.mysql.Create(data)
}

func (repo *SellingStatusRepository) Update(id int64, data *models.SellingStatus) error {
	return repo.mysql.Update(id, data)
}

func (repo *SellingStatusRepository) Delete(id int64) error {
	return repo.mysql.Delete(id)
}

func (repo *SellingStatusRepository) GetAll() ([]models.SellingStatus, error) {
	return repo.mysql.GetAll()
}

func (repo *SellingStatusRepository) GetById(id int64) (models.SellingStatus, error) {
	return repo.mysql.GetById(id)
}
