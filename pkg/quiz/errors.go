package quiz

import "fmt"

var (
	ErrorInvalidGameConfig       = fmt.Errorf("invalid game config")
	ErrorGameSessionNotFound     = fmt.Errorf("game session not found")
	ErrorInvalidStatusTransition = fmt.Errorf("invalid status transition")
)
