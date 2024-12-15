package quiz

import (
	"github.com/google/uuid"
)



func CreateGameSession(config GameSessionConfig) (*GameSession, error) {
	err := isConfigValid(config)
	if err != nil {
		return nil, err
	}

	gameSession := GameSession{
		ID:     uuid.NewString(),
		Config: config,
		Status: NotStarted,
	}

	return &gameSession, nil
}

func UpdateGameSessionConfig(id string, config GameSessionConfig) (*GameSession, error) {
	err := isConfigValid(config)
	if err != nil {
		return nil, err
	}

	for _, session := range mockGameSessions {
		if session.ID != id {
			continue
		}

		session.Config = config

		return &session, nil
	}

	return nil, ErrorGameSessionNotFound
}

func UpdateGameSessionStatus(id string, status GameStatus) (*GameSession, error) {
	for _, session := range mockGameSessions {
		if session.ID != id {
			continue
		}

		// NotStarted -> InQueue
		if session.Status == NotStarted && status == InQueue {
			session.Status = status
			return &session, nil
		}

		// InQueue -> Playing
		if session.Status == InQueue && status == Playing {
			session.Status = status
			return &session, nil
		}

		// Playing -> Finished
		if session.Status == Playing && status == Finished {
			session.Status = status
			return &session, nil
		}

		// NOTE: handle other status transitions below

		return nil, ErrorInvalidStatusTransition
	}

	return nil, ErrorGameSessionNotFound
}

func isConfigValid(config GameSessionConfig) error {
	if config.MaxPlayers < 1 {
		return ErrorInvalidGameConfig
	}

	return nil
}
