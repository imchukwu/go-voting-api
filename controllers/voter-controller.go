package controllers

import (
	"net/http"
	"go-voting-api/config"
	"go-voting-api/models"

	"github.com/gin-gonic/gin"
)

type VoterInput struct {
	FullName     string `json:"full_name"`
	SerialNumber string `json:"serial_number"`
	Class        string `json:"class"`
}

func BulkCreateVoters(c *gin.Context) {
	var input []VoterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	var voters []models.Voter
	for _, item := range input {
		voters = append(voters, models.Voter{
			FullName:     item.FullName,
			SerialNumber: item.SerialNumber,
			Class:        item.Class,
		})
	}

	if err := config.DB.Create(&voters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create voters"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voters created successfully"})
}

func RegisterVoter(c *gin.Context) {
	var input struct {
		FullName     string `json:"full_name"`
		SerialNumber string `json:"serial_number"`
		Class        string `json:"class"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	voter := models.Voter{
		FullName:     input.FullName,
		SerialNumber: input.SerialNumber,
		Class:        input.Class,
	}

	if err := config.DB.Create(&voter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register voter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voter registered successfully", "voter_id": voter.ID})
}


func VoterLogin(c *gin.Context) {
	var input struct {
		SerialNumber string `json:"serial_number"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.SerialNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Serial number is required"})
		return
	}

	var voter models.Voter
	if err := config.DB.Where("serial_number = ?", input.SerialNumber).First(&voter).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Voter not found"})
		return
	}

	// Return voter details (you could also generate a token here if needed)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login successful",
		"voter_id": voter.ID,
		"full_name": voter.FullName,
	})
}

func GetAllVoters(c *gin.Context) {
	var voters []models.Voter
	if err := config.DB.Find(&voters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch voters"})
		return
	}
	c.JSON(http.StatusOK, voters)
}

func GetVoterByID(c *gin.Context) {
	id := c.Param("id")
	var voter models.Voter
	if err := config.DB.First(&voter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voter not found"})
		return
	}
	c.JSON(http.StatusOK, voter)
}
