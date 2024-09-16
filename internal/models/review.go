package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Content        string `gorm:"type:text;not null"`
	Rating         int    `gorm:"not null;check:rating >= 1 AND rating <= 5"` // Fixed: added space between `check` and its condition
	TenderID       uint   `gorm:"not null"`
	BidID          uint   `gorm:"not null"`
	AuthorUsername string `gorm:"size:50;not null"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"` // Corrected: changed to uuid.UUID to match the schema
}
