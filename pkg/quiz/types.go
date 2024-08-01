package quiz

import "time"

type Role string

const (
	User  Role = "User"
	Admin Role = "Admin"
)

type GameSession struct {
	ID               string            `json:"id"`
	Config           GameSessionConfig `json:"config"`
	PlayerIDScoreMap map[string]int    `json:"playerIdScoreMap"`
	Status           GameStatus        `json:"status"`
}

type Stage struct {
	Question    string   `json:"question"`
	Explanation string   `json:"explanation"`
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
	BrandID       string    `json:"brandId"`
	MaxPlayers    int       `json:"maxPlayers"`
	Stages        []Stage   `json:"stages"`
	CurrentStage  int       `json:"currentStage"`
	StagePeriod   int       `json:"stagePeriod"`
	OpenQueueTime time.Time `json:"openQueueTime"`
}

type Permission string

const (
	PermissionManageGameSession Permission = "manage-game-session"
	PermissionPlayGame          Permission = "play-game"
	PermissionGetGameSession    Permission = "get-game-session"
)

type Event string

const (
	EventCreateGameSession Event = "create-game-session"
	EventUpdateGameSession Event = "update-game-session"
	EventJoinGameQueue     Event = "join-game-queue"
	EventLeaveGameQueue    Event = "leave-game-queue"
	EventGetGameSession    Event = "get-game-session"
)

type Status string

const (
	Success Status = "success"
	Failure Status = "failure"
)

type Response struct {
	Status  Status `json:"status"`
	Event   Event  `json:"event"`
	Payload any    `json:"payload"`
}
