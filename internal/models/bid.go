package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	Name            string      `gorm:"size:100;not null"`
	Description     string      `gorm:"type:text"`
	Status          string      `gorm:"size:20;default:'CREATED';not null"`
	Version         int         `gorm:"default:1"`
	TenderID        uint        `gorm:"not null"`
	Tender          Tender      `gorm:"foreignKey:TenderID"`
	OrganizationID  uuid.UUID   `gorm:"type:uuid;not null"`
	Organization    Organization `gorm:"foreignKey:OrganizationID"`
	CreatorUsername string      `gorm:"size:50;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
