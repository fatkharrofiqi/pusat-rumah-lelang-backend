package models

import "gorm.io/gorm"

type RoadAccess struct {
	gorm.Model
	IsSupportTwoRoad  bool `gorm:"is_support_two_road" json:"is_support_two_road"`
	IsSupportFourRoad bool `gorm:"is_support_four_road" json:"is_support_four_road"`
	PropertyID        uint `gorm:"property_id" json:"property_id"`
}
