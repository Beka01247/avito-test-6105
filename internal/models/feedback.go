package models

import "gorm.io/gorm"

type BidFeedback struct {
	gorm.Model
	BidID    uint   `gorm:"not null"`
	Feedback string `gorm:"type:text;not null"`
	Username string `gorm:"size:50;not null"`
}
