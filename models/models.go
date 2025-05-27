package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type Voter struct {
	gorm.Model
	FullName     string
	SerialNumber string `gorm:"unique"`
	Class        string
}

type Candidate struct {
	gorm.Model
	VoterID  uint // Link to voter
	Picture  string
	Position string

	Voter Voter `gorm:"foreignKey:VoterID"`
}

type Election struct {
	gorm.Model
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Report struct {
	gorm.Model
	ElectionID uint
	Content    string // Could be a JSON string or file path
}

// Bonus for future voting system
type Vote struct {
	gorm.Model
	VoterID     uint
	CandidateID uint
	ElectionID  uint

	Voter     Voter     `gorm:"foreignKey:VoterID"`
	Candidate Candidate `gorm:"foreignKey:CandidateID"`
	Election  Election  `gorm:"foreignKey:ElectionID"`
}
