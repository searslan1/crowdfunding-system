package model

import (
	"gorm.io/gorm"
)

// Entrepreneur girişimci modeli
type Entrepreneur struct {
	ID              uint   `gorm:"primaryKey"`
	UserID          uint   `gorm:"unique;not null"`
	StartupName     string `gorm:"not null"`
	Industry        string `gorm:"not null"`
	FundingNeeded   float64 `gorm:"not null"`
	BusinessModel   string
	PitchDeckURL     string
	Status          string `gorm:"default:'pending'"`
	CreatedAt       string `gorm:"default:CURRENT_TIMESTAMP"`
	IsEDevletApproved bool   `gorm:"default:false"` // E-devlet onayı
	IsAdminApproved  bool   `gorm:"default:false"` // Admin onayı
}
