package quiz

import "fmt"

var (
	ErrorUserNotFound            = fmt.Errorf("user not found")
	ErrorNoPermission            = fmt.Errorf("no permission")
	ErrorInvalidEvent            = fmt.Errorf("invalid event")
	ErrorInvalidGameConfig       = fmt.Errorf("invalid game config")
	ErrorGameSessionNotFound     = fmt.Errorf("game session not found")
	ErrorInvalidStatusTransition = fmt.Errorf("invalid status transition")
)
