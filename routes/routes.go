package routes

import (
	"go-voting-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Voting API is running"})
	})

	auth := r.Group("/admin")
	{
		auth.POST("/register", controllers.RegisterAdmin)
		auth.POST("/login", controllers.LoginAdmin)
	}

	voters := r.Group("/voters")
	{
		voters.POST("/register", controllers.RegisterVoter)
		voters.POST("/bulk", controllers.BulkCreateVoters)
		voters.POST("/login", controllers.VoterLogin)
		voters.GET("/", controllers.GetAllVoters)
		voters.GET("/:id", controllers.GetVoterByID)
	}

	candidates := r.Group("/candidates")
	{
		candidates.POST("/register", controllers.RegisterCandidate)
		candidates.GET("/", controllers.GetAllCandidates)
		candidates.GET("/:id", controllers.GetCandidateByID)
	}
		
	elections := r.Group("/elections")
	{
		elections.POST("/create", controllers.CreateElection)
		elections.GET("/", controllers.ListElections)
		elections.GET("/:id", controllers.GetElectionByID)
	}

	votes := r.Group("/votes")
	{
		votes.POST("/cast", controllers.CastVote)
		votes.GET("/", controllers.GetAllVotes)
    	votes.GET("/:id", controllers.GetVoteByID)
	}

	reports := r.Group("/reports")
	{
		reports.GET("/:election_id", controllers.GenerateReport)
		reports.GET("/:election_id/csv", controllers.GenerateReportCSV)
	}


}
