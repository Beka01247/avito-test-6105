package controllers

import (
	"net/http"
	"strconv"
	"zadanie-6105/internal/db"
	"zadanie-6105/internal/dto"
	"zadanie-6105/internal/helpers"
	"zadanie-6105/internal/models"

	"github.com/gin-gonic/gin"
)

func GetTenders(c *gin.Context) {
	var tenders []dto.TenderResponse  
	serviceType := c.QueryArray("service_type")

	query := db.DB.Model(&models.Tender{}).Select("id, name, description, status, service_type, version, created_at").Order("name ASC")
	if len(serviceType) > 0 {
		query = query.Where("service_type IN ?", serviceType)
	}

	if err := query.Scan(&tenders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tenders"})
		return
	}

	c.JSON(http.StatusOK, tenders)
}

func CreateTender(c *gin.Context) {
	var input models.Tender
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var employee models.Employee
	if err := db.DB.Where("username = ?", input.CreatorUsername).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !helpers.IsUserResponsibleForOrganization(input.CreatorUsername, input.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	input.Status = "CREATED"
	
	if err := db.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tender"})
		return
	}

	var response dto.TenderResponse
	query := db.DB.Model(&models.Tender{}).Select("id, name, description, status, service_type, version, created_at").Where("id = ?", input.ID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created tender"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetUserTenders(c *gin.Context) {
	
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	
	var employee models.Employee
	if err := db.DB.Where("username = ?", username).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
		return
	}

	
	var tenders []dto.TenderResponse
	query := db.DB.Model(&models.Tender{}).
		Select("id, name, description, status, service_type, version, created_at").
		Where("creator_username = ?", username).
		Limit(limit).
		Offset(offset).
		Order("name ASC")

	if err := query.Scan(&tenders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tenders"})
		return
	}

	
	c.JSON(http.StatusOK, tenders)
}


func GetTenderStatus(c *gin.Context) {
	tenderID := c.Param("tenderId")
	username := c.Query("username")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	var employee models.Employee
	if err := db.DB.Where("username = ?", username).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var tender models.Tender
	if err := db.DB.Where("id = ?", tenderID).First(&tender).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
		return
	}

	if tender.CreatorUsername != username && !helpers.IsUserResponsibleForOrganization(username, tender.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User does not have rights to view the tender status"})
		return
	}

	status := tender.Status
	var result string

	switch status {
	case "CREATED":
		result = "Created"
	case "PUBLISHED":
		result = "Published"
	case "CANCELED":
		result = "Canceled"
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": result})
}

func UpdateTenderStatus(c *gin.Context) {
	tenderID := c.Param("tenderId")
	newStatus := c.Query("status")
	username := c.Query("username")

	
	var employee models.Employee
	if err := db.DB.Where("username = ?", username).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	
	var tender models.Tender
	if err := db.DB.Where("id = ?", tenderID).First(&tender).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
		return
	}

	
	if !helpers.IsUserResponsibleForOrganization(username, tender.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	tender.Status = newStatus
	if err := db.DB.Save(&tender).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tender status"})
		return
	}

	
	var response dto.TenderResponse
	query := db.DB.Model(&models.Tender{}).Select("id, name, description, status, service_type, version, created_at").Where("id = ?", tender.ID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated tender"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func EditTender(c *gin.Context) {
	tenderID := c.Param("tenderId")
	username := c.Query("username")

	
	var employee models.Employee
	if err := db.DB.Where("username = ?", username).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	
	var tender models.Tender
	if err := db.DB.Where("id = ?", tenderID).First(&tender).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
		return
	}

	
	if !helpers.IsUserResponsibleForOrganization(username, tender.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	tenderVersion := models.TenderVersion{
		TenderID:      tender.ID,
		VersionNumber: tender.Version,
		Name:          tender.Name,
		Description:   tender.Description,
		ServiceType:   tender.ServiceType,
		Status:        tender.Status,
		UpdatedBy:     username,
		UpdatedAt:     tender.UpdatedAt,
	}
	if err := db.DB.Create(&tenderVersion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save tender version"})
		return
	}

	
	if err := db.DB.Model(&tender).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tender"})
		return
	}

	
	tender.Version += 1
	db.DB.Save(&tender)

	
	var response dto.TenderResponse
	query := db.DB.Model(&models.Tender{}).Select("id, name, description, status, service_type, version, created_at").Where("id = ?", tender.ID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated tender"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func RollbackTender(c *gin.Context) {
	tenderID := c.Param("tenderId")
	versionStr := c.Param("version")
	username := c.Query("username")

	
	if tenderID == "" || versionStr == "" || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tender ID, version, and username are required"})
		return
	}

	
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid version number"})
		return
	}

	
	var employee models.Employee
	if err := db.DB.Where("username = ?", username).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	
	var tender models.Tender
	if err := db.DB.Where("id = ?", tenderID).First(&tender).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
		return
	}

	
	if !helpers.IsUserResponsibleForOrganization(username, tender.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	var tenderVersion models.TenderVersion
	if err := db.DB.Where("tender_id = ? AND version_number = ?", tenderID, version).First(&tenderVersion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Version not found"})
		return
	}

	
	tender.Name = tenderVersion.Name
	tender.Description = tenderVersion.Description
	tender.ServiceType = tenderVersion.ServiceType
	tender.Status = tenderVersion.Status
	tender.Version = tenderVersion.VersionNumber 

	
	if err := db.DB.Save(&tender).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rollback tender"})
		return
	}

	
	var response dto.TenderResponse
	query := db.DB.Model(&models.Tender{}).Select("id, name, description, status, service_type, version, created_at").Where("id = ?", tender.ID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve rolled-back tender"})
		return
	}

	c.JSON(http.StatusOK, response)
}