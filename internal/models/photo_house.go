package models

import "gorm.io/gorm"

type PhotoHouse struct {
	gorm.Model
	PropertyID uint   `gorm:"property_id" json:"property_id"`
	PhotoUrl   string `gorm:"photo_url" json:"photo_url"`
}
