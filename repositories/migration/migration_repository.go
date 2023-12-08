package migration

import (
	"pusat-rumah-lelang-backend/migrations"
	"pusat-rumah-lelang-backend/seeder"

	"gorm.io/gorm"
)

type IMigrationRepository interface {
	Migrate()
}

type MigrationRepository struct {
	db *gorm.DB
}

func NewMigrationRepository(db *gorm.DB) IMigrationRepository {
	return &MigrationRepository{db: db}
}

func (m *MigrationRepository) Migrate() {
	migrations.RunMigrations(m.db)

	seeder.Seeds(m.db)
}
