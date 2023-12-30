package seeder

import (
	"pusat-rumah-lelang-backend/internal/model"

	"gorm.io/gorm"
)

func Seeds(db *gorm.DB) {
	roadAccess := []model.RoadAccess{
		{Name: "Roda 4", Description: "Roda 4"},
		{Name: "Roda 2", Description: "Roda 2"},
	}
	if err := db.Create(&roadAccess).Error; err != nil {
		panic(err.Error)
	}

	bank := []model.Bank{
		{Name: "BPR Bintara", Description: "BPR Bintara"},
	}

	if err := db.Create(&bank).Error; err != nil {
		panic(err.Error)
	}

	sellingStatus := []model.SellingStatus{
		{Name: "Suka Rela", Description: "Description for Suka Rela"},
		{Name: "Lelang", Description: "Description for Lelang"},
		{Name: "Ayda", Description: "Description for Ayda"},
		{Name: "Rumah Biasa", Description: "Description for Rumah Biasa"},
	}

	if err := db.Create(&sellingStatus).Error; err != nil {
		panic(err.Error)
	}

	// Create a Property instance
	property := model.Property{
		Address:             "123 Main St",
		BuildingArea:        "110 m",
		LandArea:            "200 m",
		ElectricityCapacity: "1300 VA",
		WaterSource:         "Jetpam Sanyo",
		Owner:               "Arba",
		Title:               "Rumah Bekasi murah dan mantaps",
		Bedrooms:            3,
		Price:               1500000000.0,
		Latitude:            "0.3234293",
		Longitude:           "02939283",
		BankID:              &bank[0].ID,
		SellingStatusID:     &sellingStatus[0].ID,
		PropertyTaxPhoto:    "https://placekitten.com/g/500/300",
		Description:         "A beautiful property for sale",
		RoadAccessID:        &roadAccess[0].ID,
		PhotoHouse: []*model.PhotoHouse{
			{PhotoUrl: "https://placekitten.com/g/500/300"},
			{PhotoUrl: "https://placekitten.com/g/600/300"},
		},
		PhotoCertificate: []*model.PhotoCertificate{
			{CertificateUrl: "https://placekitten.com/g/500/300"},
			{CertificateUrl: "https://placekitten.com/g/400/200"},
		},
	}

	// Create the Property record and its related PhotoHouse and PhotoCertificate records
	if err := db.Create(&property).Error; err != nil {
		panic(err)
	}
}
