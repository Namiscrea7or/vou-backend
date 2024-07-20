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
	Question    string   `json:"question"`
	Options     []string `json:"options"`
	AnswerIndex int      `json:"answerIndex"`
	Points      int      `json:"points"`
}

type GameStatus string

const (
	NotStarted GameStatus = "Not Started"
	InQueue    GameStatus = "In Queue"
	Playing    GameStatus = "Playing"
	Finished   GameStatus = "Finished"
)

type GameSessionConfig struct {
	BrandID       string  `json:"brandId"`
	MaxPlayers    int     `json:"maxPlayers"`
	Stages        []Stage `json:"stages"`
	CurrentStage  int
	StagePeriod   int
	OpenQueueTime time.Time
	StartTime     time.Time
}

type Permission string

const (
	PermissionManageGameSession Permission = "manage-game-session"
	PermissionPlayGame          Permission = "play-game"
)

type Event string

const (
	EventCreateGameSession       Event = "create-game-session"
	EventUpdateGameSessionConfig Event = "update-game-session-config"
	EventJoinGameQueue           Event = "join-game-queue"
	EventLeaveGameQueue          Event = "leave-game-queue"
)

type Status string

const (
	Success Status = "success"
	Failure Status = "failure"
)

type Response struct {
	Type    Status `json:"type"`
	Event   Event  `json:"event"`
	Payload any    `json:"payload"`
}
