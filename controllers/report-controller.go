package controllers

import (
	"encoding/csv"
	"go-voting-api/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CandidateReport struct {
	CandidateID uint   `json:"candidate_id"`
	FullName    string `json:"full_name"`
	Position    string `json:"position"`
	Votes       int64  `json:"votes"`
}

func GenerateReport(c *gin.Context) {
	electionID := c.Param("election_id")

	var results []CandidateReport

	err := config.DB.Table("votes").
		Select("candidates.id AS candidate_id, voters.full_name, candidates.position, COUNT(votes.id) AS votes").
		Joins("JOIN candidates ON candidates.id = votes.candidate_id").
		Joins("JOIN voters ON voters.id = candidates.voter_id").
		Where("votes.election_id = ?", electionID).
		Group("candidates.id, voters.full_name, candidates.position").
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"election_id": electionID,
		"results":     results,
	})
}

func GenerateReportCSV(c *gin.Context) {
	electionID := c.Param("election_id")

	var results []struct {
		CandidateID uint
		FullName    string
		Position    string
		Votes       int64
	}

	err := config.DB.Table("votes").
		Select("candidates.id AS candidate_id, voters.full_name, candidates.position, COUNT(votes.id) AS votes").
		Joins("JOIN candidates ON candidates.id = votes.candidate_id").
		Joins("JOIN voters ON voters.id = candidates.voter_id").
		Where("votes.election_id = ?", electionID).
		Group("candidates.id, voters.full_name, candidates.position").
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate report"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=report_election_"+electionID+".csv")
	c.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Candidate ID", "Full Name", "Position", "Votes"})

	// Write CSV rows
	for _, r := range results {
		writer.Write([]string{
			strconv.FormatUint(uint64(r.CandidateID), 10),
			r.FullName,
			r.Position,
			strconv.FormatInt(r.Votes, 10),
		})
	}
}
