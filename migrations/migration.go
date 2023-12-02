package migrations

import (
	"pusat-rumah-lelang-backend/internal/models"

	"gorm.io/gorm"
)

// migrates table.
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.SellingStatus{},
		&models.Property{},
		&models.PhotoHouse{},
		&models.PhotoCertificate{},
	)
}
