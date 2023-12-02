package models

import "gorm.io/gorm"

type PhotoCertificate struct {
	gorm.Model
	PropertyID     uint   `gorm:"property_id" json:"property_id"`
	CertificateUrl string `gorm:"certificate_url" json:"certificate_url"`
}
