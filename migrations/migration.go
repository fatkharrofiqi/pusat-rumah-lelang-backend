package migrations

import (
	"pusat-rumah-lelang-backend/models"

	"gorm.io/gorm"
)

// migrates table.
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Certificate{},
		&models.Bank{},
		&models.SellingStatus{},
		&models.Property{},
		&models.RoadAccess{},
		&models.PhotoHouse{},
		&models.PhotoCertificate{},
	)
}

func DropTable(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.Certificate{},
		&models.Bank{},
		&models.SellingStatus{},
		&models.Property{},
		&models.RoadAccess{},
		&models.PhotoHouse{},
		&models.PhotoCertificate{})
}
