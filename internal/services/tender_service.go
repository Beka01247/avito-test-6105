package services

import (
	"database/sql"
	"log"
	"zadanie-6105/internal/models"
)

func CreateTender(db *sql.DB, tender models.Tender) error {
	query := `INSERT INTO tenders (name, description, service_type, status, organization_id, version)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationID, 1)
	if err != nil {
		log.Printf("Error creating tender: %v", err)
		return err
	}
	return nil
}



















