package controllers

import (
	"net/http"
	"time"
	"go-voting-api/config"
	"go-voting-api/models"

	"github.com/gin-gonic/gin"
)

func CreateElection(c *gin.Context) {
	var input struct {
		Title     string `json:"title"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	start, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}

	end, err := time.Parse(time.RFC3339, input.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}

	election := models.Election{
		Title:     input.Title,
		StartTime: start,
		EndTime:   end,
	}

	if err := config.DB.Create(&election).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create election"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Election created successfully"})
}

func ListElections(c *gin.Context) {
	var elections []models.Election
	if err := config.DB.Find(&elections).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve elections"})
		return
	}
	c.JSON(http.StatusOK, elections)
}

func GetElectionByID(c *gin.Context) {
	id := c.Param("id")
	var election models.Election
	if err := config.DB.First(&election, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Election not found"})
		return
	}
	c.JSON(http.StatusOK, election)
}
