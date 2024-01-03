package model

import (
	"time"

	"gorm.io/gorm"
)

type PhotoHouse struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	PropertyID  uint           `gorm:"property_id" json:"property_id"`
	PhotoUrl    string         `gorm:"photo_url" json:"photo_url"`
	IsThumbnail bool           `gorm:"is_thumbnail" json:"is_thumbnail"`
	IsCover     bool           `gorm:"is_cover" json:"is_cover"`
}
