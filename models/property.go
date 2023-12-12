package models

import (
	"time"

	"gorm.io/gorm"
)

type Property struct {
	ID                  uint                `gorm:"primarykey" json:"id"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
	DeletedAt           gorm.DeletedAt      `gorm:"index" json:"deleted_at"`
	Title               string              `grom:"title" json:"title"`
	Owner               string              `gorm:"owner" json:"owner"`
	Address             string              `gorm:"address" json:"address"`
	BuildingArea        string              `gorm:"building_area" json:"building_area"`
	LandArea            string              `gorm:"land_area" json:"land_area"`
	Latitude            string              `gorm:"latitude" json:"latitude"`
	Longitude           string              `gorm:"longitude" json:"longitude"`
	PropertyTaxPhoto    string              `gorm:"property_tax_photo" json:"property_tax"`
	ElectricityCapacity string              `gorm:"electricity_capacity" json:"electricity_capacity"`
	WaterSource         string              `gorm:"water_source" json:"water_source"`
	Bedrooms            int                 `gorm:"bedrooms" json:"bedrooms"`
	Price               float64             `gorm:"price" json:"price"`
	SellingStatusID     *uint               `gorm:"selling_status_id" json:"selling_status_id"`
	SellingStatus       *SellingStatus      `json:"selling_status"`
	BankID              *uint               `gorm:"bank_id" json:"bank_id"`
	Bank                *Bank               `json:"bank"`
	RoadAccessID        *uint               `gorm:"road_access_id;" json:"road_access_id"`
	RoadAccess          *RoadAccess         `json:"road_access"`
	Description         string              `gorm:"description" json:"description"`
	PhotoHouse          []*PhotoHouse       `gorm:"photo_house" json:"photo_houses"`
	PhotoCertificate    []*PhotoCertificate `gorm:"photo_certificate" json:"photo_certificates"`
	CertificateID       *uint               `gorm:"certificate_id" json:"certificate_id"`
	Certificate         *Certificate        `json:"certificate"`
}
