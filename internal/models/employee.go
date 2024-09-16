package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // UUID type and correct default for GORM
	Username  string    `gorm:"size:50;unique;not null"`
	FirstName string    `gorm:"size:50"`
	LastName  string    `gorm:"size:50"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Match the database default
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Match the database default
}