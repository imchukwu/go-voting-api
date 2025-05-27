package controllers

import (
	"net/http"
	"go-voting-api/config"
	"go-voting-api/models"

	"github.com/gin-gonic/gin"
)

func RegisterCandidate(c *gin.Context) {
	var input struct {
		VoterID  uint   `json:"voter_id"`
		Picture  string `json:"picture"`
		Position string `json:"position"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Optional: check that voter exists
	var voter models.Voter
	if err := config.DB.First(&voter, input.VoterID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Voter not found"})
		return
	}

	candidate := models.Candidate{
		VoterID:  input.VoterID,
		Picture:  input.Picture,
		Position: input.Position,
	}

	if err := config.DB.Create(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register candidate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Candidate registered successfully"})
}

func GetAllCandidates(c *gin.Context) {
	var candidates []models.Candidate
	if err := config.DB.Preload("Voter").Find(&candidates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch candidates"})
		return
	}
	c.JSON(http.StatusOK, candidates)
}

func GetCandidateByID(c *gin.Context) {
	id := c.Param("id")
	var candidate models.Candidate
	if err := config.DB.Preload("Voter").First(&candidate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}
	c.JSON(http.StatusOK, candidate)
}
