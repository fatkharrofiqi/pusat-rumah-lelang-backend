package models

import "gorm.io/gorm"

type RoadAccess struct {
	gorm.Model
	Name        string      `gorm:"name" json:"name"`
	Description string      `gorm:"description" json:"description"`
	Property    []*Property `gorm:"many2many:property_road" json:"properties"`
}
