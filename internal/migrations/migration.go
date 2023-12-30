package migrations

import (
	"pusat-rumah-lelang-backend/internal/model"

	"gorm.io/gorm"
)

// migrates table.
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Certificate{},
		&model.Bank{},
		&model.SellingStatus{},
		&model.Property{},
		&model.RoadAccess{},
		&model.PhotoHouse{},
		&model.PhotoCertificate{},
	)
}

func DropTable(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&model.Certificate{},
		&model.Bank{},
		&model.SellingStatus{},
		&model.Property{},
		&model.RoadAccess{},
		&model.PhotoHouse{},
		&model.PhotoCertificate{})
}
