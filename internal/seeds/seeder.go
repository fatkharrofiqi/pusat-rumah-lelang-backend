package seeds

import (
	"pusat-rumah-lelang-backend/internal/models"

	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {

	bank := []models.Bank{
		{Name: "Bank Bintara", Description: "Bank Bintara"},
		{Name: "Bank ABC", Description: "Bank Bintara"},
	}

	if err := db.Create(&bank).Error; err != nil {
		panic(err.Error)
	}

	sellingStatus := []models.SellingStatus{
		{Name: "Penjualan Sukarela", Description: "Description for Penjualan Sukarela"},
		{Name: "Lelang", Description: "Description for Lelang"},
		{Name: "Ayda", Description: "Description for Ayda"},
		{Name: "Jual Rumah Biasa", Description: "Description for Jual Rumah Biasa"},
	}

	if err := db.Create(&sellingStatus).Error; err != nil {
		panic(err.Error)
	}

	firstSellingStatus := models.SellingStatus{}

	db.First(&firstSellingStatus)

	// Create a Property instance
	property := models.Property{
		Address:             "123 Main St",
		BuildingArea:        "110 m",
		LandArea:            "200 m",
		ElectricityCapacity: 1300,
		WaterSource:         "Jetpam Sanyo",
		Owner:               "Arba",
		Title:               "Rumah Bekasi murah dan mantaps",
		Bedrooms:            3,
		Price:               1500000000.0,
		Latitude:            "0.3234293",
		Longitude:           "02939283",
		BankID:              &bank[0].ID,
		SellingStatusID:     &firstSellingStatus.ID,
		PropertyTaxPhoto:    "https://placekitten.com/g/500/300",
		Description:         "A beautiful property for sale",
		RoadAccess: models.RoadAccess{
			IsSupportTwoRoad:  true,
			IsSupportFourRoad: true,
		},
		PhotoHouse: []*models.PhotoHouse{
			{PhotoUrl: "https://placekitten.com/g/500/300"},
			{PhotoUrl: "https://placekitten.com/g/600/300"},
		},
		PhotoCertificate: []*models.PhotoCertificate{
			{CertificateUrl: "https://placekitten.com/g/500/300"},
			{CertificateUrl: "https://placekitten.com/g/400/200"},
		},
	}

	// Create the Property record and its related PhotoHouse and PhotoCertificate records
	if err := db.Create(&property).Error; err != nil {
		panic(err)
	}
}
