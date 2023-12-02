package models

import "gorm.io/gorm"

type SellingStatus struct {
	gorm.Model
	Name        string `gorm:"name" json:"name"`
	Description string `gorm:"description" json:"description"`
}
