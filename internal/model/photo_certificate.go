package model

import (
	"time"

	"gorm.io/gorm"
)

type PhotoCertificate struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	PropertyID     uint           `gorm:"property_id" json:"property_id"`
	CertificateUrl string         `gorm:"certificate_url" json:"certificate_url"`
}
