package usecases

import "pusat-rumah-lelang-backend/repositories/migration"

type IMigrationUsecase interface {
	Migrate()
}

type MigrationUsecase struct {
	repo migration.IMigrationRepository
}

func NewMigrationUsecase(repo migration.IMigrationRepository) IMigrationUsecase {
	return &MigrationUsecase{repo: repo}
}

func (u *MigrationUsecase) Migrate() {
	u.repo.Migrate()
}
