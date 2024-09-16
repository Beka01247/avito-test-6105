package models

import (
	"time"

	"gorm.io/gorm"
)

type BidVersion struct {
	gorm.Model
	BidID         uint      `gorm:"not null"`
	VersionNumber int       `gorm:"not null"`
	Name          string    `gorm:"size:100"`
	Description   string    `gorm:"type:text"`
	Status        string    `gorm:"size:50"`
	UpdatedBy     string    `gorm:"size:50"`
	UpdatedAt     time.Time
}
