package models

import (
	"time"

	"gorm.io/gorm"
)

type SellingStatus struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `gorm:"name" json:"name"`
	Description string         `gorm:"description" json:"description"`
	Property    []*Property    `gorm:"property" json:"properties,omitempty"`
}
