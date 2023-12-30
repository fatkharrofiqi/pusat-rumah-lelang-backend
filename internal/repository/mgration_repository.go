package repository

import (
	"pusat-rumah-lelang-backend/db/seeder"
	"pusat-rumah-lelang-backend/internal/migrations"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IMigrationRepository interface {
	Migrate(db *gorm.DB)
}

type MigrationRepository struct {
	Log *logrus.Logger
}

func NewMigrationRepository(log *logrus.Logger) IMigrationRepository {
	return &MigrationRepository{Log: log}
}

func (m *MigrationRepository) Migrate(db *gorm.DB) {
	migrations.RunMigrations(db)

	seeder.Seeds(db)
}
