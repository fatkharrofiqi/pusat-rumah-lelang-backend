package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Address             string             `gorm:"address" json:"address"`
	BuildingArea        string             `gorm:"building_area" json:"building_area"`
	LandArea            string             `gorm:"land_area" json:"land_area"`
	Latitude            string             `gorm:"latitude" json:"latitude"`
	Longitude           string             `gorm:"longitude" json:"longitude"`
	PropertyTaxPhoto    string             `gorm:"property_tax_photo" json:"property_tax"`
	ElectricityCapacity int                `gorm:"electricity_capacity" json:"electricity_capacity"`
	WaterSource         string             `gorm:"water_source" json:"water_source"`
	RoadAccess          string             `gorm:"road_access" json:"road_access"`
	Bedrooms            int                `gorm:"bedrooms" json:"bedrooms"`
	Price               float64            `gorm:"price" json:"price"`
	SellingStatusID     *uint              `gorm:"selling_status_id" json:"selling_status_id"`
	SellingStatus       *SellingStatus     `json:"selling_status"`
	Description         string             `gorm:"description" json:"description"`
	PhotoHouse          []PhotoHouse       `gorm:"photo_house" json:"photo_houses"`
	PhotoCertificate    []PhotoCertificate `gorm:"photo_certificate" json:"photo_certificates"`
}
