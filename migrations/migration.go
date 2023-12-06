package migrations

import (
	"pusat-rumah-lelang-backend/models"

	"gorm.io/gorm"
)

// migrates table.
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Bank{},
		&models.SellingStatus{},
		&models.Property{},
		&models.RoadAccess{},
		&models.PhotoHouse{},
		&models.PhotoCertificate{},
	)
}
