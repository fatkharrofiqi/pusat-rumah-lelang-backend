package usecase

import (
	"pusat-rumah-lelang-backend/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IMigrationUsecase interface {
	Migrate(c *gin.Context)
}

type MigrationUsecase struct {
	MigratonRepository repository.IMigrationRepository
	DB                 *gorm.DB
}

func NewMigrationUsecase(db *gorm.DB, repo repository.IMigrationRepository) IMigrationUsecase {
	return &MigrationUsecase{MigratonRepository: repo, DB: db}
}

func (u *MigrationUsecase) Migrate(c *gin.Context) {
	u.MigratonRepository.Migrate(u.DB.WithContext(c))
}
