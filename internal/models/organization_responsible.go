package models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationResponsible struct {
	ID            uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID type and correct default for GORM
	OrganizationID uuid.UUID  `gorm:"type:uuid;not null"`
	Organization  Organization `gorm:"foreignKey:OrganizationID"`
	UserID        uuid.UUID   `gorm:"type:uuid;not null"`
	Employee      Employee    `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP"`
}