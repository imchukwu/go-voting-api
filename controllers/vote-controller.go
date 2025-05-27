package controllers

import (
	"net/http"
	"go-voting-api/config"
	"go-voting-api/models"

	"github.com/gin-gonic/gin"
)

func CastVote(c *gin.Context) {
	var input struct {
		VoterID     uint `json:"voter_id"`
		CandidateID uint `json:"candidate_id"`
		ElectionID  uint `json:"election_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get candidate to check position
	var candidate models.Candidate
	if err := config.DB.First(&candidate, input.CandidateID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}

	// Check if voter already voted for this position in this election
	var existingVote models.Vote
	err := config.DB.
		Joins("JOIN candidates ON candidates.id = votes.candidate_id").
		Where("votes.voter_id = ? AND votes.election_id = ? AND candidates.position = ?", input.VoterID, input.ElectionID, candidate.Position).
		First(&existingVote).Error

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "You already voted for this position"})
		return
	}

	vote := models.Vote{
		VoterID:     input.VoterID,
		CandidateID: input.CandidateID,
		ElectionID:  input.ElectionID,
	}

	if err := config.DB.Create(&vote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Vote not recorded"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote cast successfully"})
}

func GetAllVotes(c *gin.Context) {
	var votes []models.Vote
	if err := config.DB.Preload("Candidate").Preload("Voter").Preload("Election").Find(&votes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch votes"})
		return
	}
	c.JSON(http.StatusOK, votes)
}

func GetVoteByID(c *gin.Context) {
	id := c.Param("id")
	var vote models.Vote
	if err := config.DB.Preload("Candidate").Preload("Voter").Preload("Election").First(&vote, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vote not found"})
		return
	}
	c.JSON(http.StatusOK, vote)
}
