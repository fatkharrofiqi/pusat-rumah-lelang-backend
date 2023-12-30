package model

import (
	"time"

	"gorm.io/gorm"
)

type Certificate struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
}
