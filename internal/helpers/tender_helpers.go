package helpers

import (
	"zadanie-6105/internal/db"
	"zadanie-6105/internal/models"

	"github.com/google/uuid"
)

func IsUserResponsibleForOrganization(username string, organizationID uuid.UUID) bool {
	var responsible models.OrganizationResponsible
	if err := db.DB.Where("user_id = (SELECT id FROM employee WHERE username = ?) AND organization_id = ?", username, organizationID).First(&responsible).Error; err != nil {
		return false
	}
	return true
}

func CheckTenderExists(tenderID uint) bool {
	var tender models.Tender
	if err := db.DB.Where("id = ?", tenderID).First(&tender).Error; err != nil {
		return false
	}
	return true
}

func CheckBidExists(bidID uint) bool {
	var bid models.Bid
	if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
		return false
	}
	return true
}
