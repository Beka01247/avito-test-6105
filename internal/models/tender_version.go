package models

import (
	"time"

	"gorm.io/gorm"
)

type TenderVersion struct {
	gorm.Model
	TenderID      uint      `gorm:"not null"`
	VersionNumber int       `gorm:"not null"`
	Name          string    `gorm:"size:100"`
	Description   string    `gorm:"type:text"`
	ServiceType   string    `gorm:"size:50"`
	Status        string    `gorm:"size:50"`
	UpdatedBy     string    `gorm:"size:50"`
	UpdatedAt     time.Time
}
