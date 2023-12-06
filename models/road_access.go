package models

import (
	"time"

	"gorm.io/gorm"
)

type RoadAccess struct {
	ID                uint           `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	IsSupportTwoRoad  bool           `gorm:"is_support_two_road" json:"is_support_two_road"`
	IsSupportFourRoad bool           `gorm:"is_support_four_road" json:"is_support_four_road"`
	PropertyID        uint           `gorm:"property_id" json:"property_id"`
}
