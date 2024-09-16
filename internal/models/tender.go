package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tender struct {
	gorm.Model
	Name            string    `gorm:"size:100;not null"`
	Description     string    `gorm:"type:text"`
	ServiceType     string    `gorm:"size:50;not null"`
	Status          string    `gorm:"size:20;default:'CREATED';not null"`
	OrganizationID  uuid.UUID `gorm:"type:uuid;not null"`
	CreatorUsername string    `gorm:"size:50;not null"`
	Version         int       `gorm:"default:1"`
}
