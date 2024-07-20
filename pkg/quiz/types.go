package quiz

import "time"

type Role string

const (
	User  Role = "User"
	Admin Role = "Admin"
)

type GameSession struct {
	ID               string
	Config           GameSessionConfig
	PlayerIDScoreMap map[string]int
	Status           GameStatus
}

type Stage struct {
	Question    string
	Options     []string
	AnswerIndex int
	Points      int
}

type GameStatus string

const (
	NotStarted GameStatus = "Not Started"
	InQueue    GameStatus = "In Queue"
	Playing    GameStatus = "Playing"
	Finished   GameStatus = "Finished"
)

type GameSessionConfig struct {
	BrandID       string
	MaxPlayers    int
	Stages        []Stage
	OpenQueueTime time.Time
	StartTime     time.Time
}
