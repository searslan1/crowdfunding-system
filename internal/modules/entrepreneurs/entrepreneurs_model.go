package model

import (
	"time"

	"gorm.io/gorm"
)

type EntrepreneurStatus string

const (
	Pending  EntrepreneurStatus = "pending"
	Approved EntrepreneurStatus = "approved"
	Rejected EntrepreneurStatus = "rejected"
)

type Entrepreneur struct {
	ID            uint                `gorm:"primaryKey" json:"id"`
	UserID        uint                `gorm:"unique;not null" json:"user_id"`
	StartupName   string              `gorm:"type:varchar(255);not null" json:"startup_name"`
	Industry      string              `gorm:"type:varchar(100);not null" json:"industry"`
	FundingNeeded float64             `gorm:"type:decimal(15,2);not null" json:"funding_needed"`
	BusinessModel *string             `gorm:"type:text" json:"business_model,omitempty"`
	PitchDeckURL  *string             `gorm:"type:text" json:"pitch_deck_url,omitempty"`
	Status        EntrepreneurStatus   `gorm:"type:entrepreneur_status;default:'pending'" json:"status"`
	CreatedAt     time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt      `gorm:"index" json:"-"`
}
