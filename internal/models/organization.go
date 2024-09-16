package models

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID type and correct default for GORM
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	Type        string    `gorm:"type:organization_type;not null"` // Enum type
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`       // Match the database default
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`       // Match the database default
	Tenders     []Tender  `gorm:"foreignKey:OrganizationID"`
	Bids        []Bid     `gorm:"foreignKey:OrganizationID"`
	Reviews     []Review  `gorm:"foreignKey:OrganizationID"`
}

