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

func CreateBid(c *gin.Context) {
	var input models.Bid
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var employee models.Employee
	if err := db.DB.Where("username = ?", input.CreatorUsername).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var organization models.Organization
	if err := db.DB.Where("id = ?", input.OrganizationID).First(&organization).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	var tender models.Tender
	if err := db.DB.Where("id = ?", input.TenderID).First(&tender).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
		return
	}

	if !helpers.IsUserResponsibleForOrganization(input.CreatorUsername, input.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	if err := db.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bid"})
		return
	}

	
	var response dto.BidResponse
	query := db.DB.Model(&models.Bid{}).Select("id, name, status, version, created_at").Where("id = ?", input.ID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve created bid"})
		return
	}

	
	response.AuthorType = "User" 
	response.AuthorID = employee.ID

	c.JSON(http.StatusOK, response)
}


func GetUserBids(c *gin.Context) {
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

	
	var bids []dto.BidResponse
	query := db.DB.Model(&models.Bid{}).Select("id, name, status, version, created_at").Where("creator_username = ?", username).Limit(limit).Offset(offset).Order("name ASC")

	if err := query.Scan(&bids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bids"})
		return
	}

	
	for i := range bids {
		bids[i].AuthorType = "User"
		bids[i].AuthorID = employee.ID
	}

	c.JSON(http.StatusOK, bids)
}


func GetBidsForTender(c *gin.Context) {
	tenderID := c.Param("tenderId")

	
	var bids []models.Bid
	if err := db.DB.Where("tender_id = ?", tenderID).Order("name ASC").Find(&bids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bids"})
		return
	}

	c.JSON(http.StatusOK, bids)
}


func GetBidStatus(c *gin.Context) {
	bidID := c.Param("bidId")
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

	
	var bid models.Bid
	if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bid not found"})
		return
	}

	
	if !helpers.IsUserResponsibleForOrganization(username, bid.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"status": bid.Status})
}


func UpdateBidStatus(c *gin.Context) {
	bidID := c.Param("bidId")
	newStatus := c.Query("status")
	username := c.Query("username")

	
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	
	if newStatus != "Created" && newStatus != "Published" && newStatus != "Canceled" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
		return
	}	

	
	var employee models.Employee
	if err := db.DB.Where("username = ?", username).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	
	var bid models.Bid
	if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bid not found"})
		return
	}

	
	if !helpers.IsUserResponsibleForOrganization(username, bid.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	bid.Status = newStatus
	if err := db.DB.Save(&bid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bid status"})
		return
	}

	
	var response dto.BidResponse
	query := db.DB.Model(&models.Bid{}).Select("id, name, status, version, created_at").Where("id = ?", bidID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated bid"})
		return
	}

	
	response.AuthorType = "User"
	response.AuthorID = employee.ID

	c.JSON(http.StatusOK, response)
}

func EditBid(c *gin.Context) {
	bidID := c.Param("bidId")
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

	
	var bid models.Bid
	if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bid not found"})
		return
	}

	
	if !helpers.IsUserResponsibleForOrganization(username, bid.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	bidVersion := models.BidVersion{
		BidID:         bid.ID,
		VersionNumber: bid.Version,
		Name:          bid.Name,
		Description:   bid.Description,
		Status:        bid.Status,
		UpdatedBy:     username,
		UpdatedAt:     bid.UpdatedAt,
	}
	if err := db.DB.Create(&bidVersion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save bid version"})
		return
	}

	
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	
	if err := db.DB.Model(&bid).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bid"})
		return
	}

	
	bid.Version += 1
	db.DB.Save(&bid)

	
	var response dto.BidResponse
	query := db.DB.Model(&models.Bid{}).Select("id, name, status, version, created_at").Where("id = ?", bidID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated bid"})
		return
	}

	
	response.AuthorType = "User"
	response.AuthorID = employee.ID

	c.JSON(http.StatusOK, response)
}



func SubmitBidDecision(c *gin.Context) {
	bidID := c.Param("bidId")
	decision := c.Query("decision")
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

	
	var bid models.Bid
	if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bid not found"})
		return
	}

	
	if !helpers.IsUserResponsibleForOrganization(username, bid.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	if decision == "Approved" {
		bid.Status = "Approved"
	} else if decision == "Rejected" {
		bid.Status = "Rejected"
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid decision"})
		return
	}

	
	if err := db.DB.Save(&bid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit bid decision"})
		return
	}

	
	var response dto.BidResponse
	query := db.DB.Model(&models.Bid{}).Select("id, name, status, version, created_at").Where("id = ?", bidID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated bid"})
		return
	}

	
	response.AuthorType = "User"
	response.AuthorID = employee.ID

	c.JSON(http.StatusOK, response)
}


func RollbackBid(c *gin.Context) {
	bidID := c.Param("bidId")
	versionStr := c.Param("version")
	username := c.Query("username")

	
	if username == "" || versionStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and version are required"})
		return
	}

	
	var employee models.Employee
	if err := db.DB.Where("username = ?", username).First(&employee).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid version number"})
		return
	}

	
	var bid models.Bid
	if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bid not found"})
		return
	}

	
	if !helpers.IsUserResponsibleForOrganization(username, bid.OrganizationID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not responsible for this organization"})
		return
	}

	
	var bidVersion models.BidVersion
	if err := db.DB.Where("bid_id = ? AND version_number = ?", bidID, version).First(&bidVersion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Version not found"})
		return
	}

	
	bid.Name = bidVersion.Name
	bid.Description = bidVersion.Description
	bid.Status = bidVersion.Status
	bid.Version = version 

	
	if err := db.DB.Save(&bid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bid"})
		return
	}

	
	var response dto.BidResponse
	query := db.DB.Model(&models.Bid{}).Select("id, name, status, version, created_at").Where("id = ?", bidID)

	if err := query.Scan(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated bid"})
		return
	}

	
	response.AuthorType = "User"
	response.AuthorID = employee.ID

	c.JSON(http.StatusOK, response)
}
