package controllers

import (
	"net/http"
	"strconv"
	"zadanie-6105/internal/db"
	"zadanie-6105/internal/models"

	"github.com/gin-gonic/gin"
)

func SubmitBidFeedback(c *gin.Context) {
	bidID := c.Param("bidId")
	feedback := c.Query("bidFeedback")
	username := c.Query("username")

	if feedback == "" || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feedback and username are required"})
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

	
	bidFeedback := models.BidFeedback{
		BidID:    bid.ID,
		Feedback: feedback,
		Username: username,
	}

	if err := db.DB.Create(&bidFeedback).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit feedback"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"feedback": bidFeedback})
}

func GetBidReviews(c *gin.Context) {
	tenderID := c.Param("tenderId")
	authorUsername := c.Query("authorUsername")
	requesterUsername := c.Query("requesterUsername")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	if authorUsername == "" || requesterUsername == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author and requester usernames are required"})
		return
	}

	
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

	
	

	
	var bids []models.Bid
	if err := db.DB.Where("creator_username = ? AND tender_id = ?", authorUsername, tenderID).Find(&bids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bids"})
		return
	}

	
	var feedbacks []models.BidFeedback
	if err := db.DB.Where("bid_id IN ?", bids).Limit(limit).Offset(offset).Find(&feedbacks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reviews"})
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}
