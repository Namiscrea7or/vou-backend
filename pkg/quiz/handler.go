package quiz

import (
	"context"
	"encoding/json"
	"fmt"

	"vou/pkg/ws"
)

func HandleGameEvent(ctx context.Context, message ws.Message) (*Response, error) {
	// Step 1: Get jwt from context
	idToken := ctx.Value(ws.AuthKey).(string)

	// Step 2: Get user profile based on jwt
	userID := idToken // need to handle correctly

	// Step 3: Authorize user based on his/her role and event
	role, found := mockUserIDRoleMap[userID]
	if !found {
		// NOTE: handle this logic based on biz requirements, I'll make it User for now.
		role = User
	}

	permission, err := getPermissionByEvent(message.Event)
	if err != nil {
		return nil, err
	}

	if !hasPermission(role, permission) {
		return nil, ErrorNoPermission
	}

	// Step 4: Execute event
	switch message.Event {
	case string(EventCreateGameSession):
		return nil, nil

	case string(EventGetGameSession):
		gameSessionId, found := message.Payload.(map[string]interface{})["sessionId"].(string)
		if !found {
			return nil, fmt.Errorf("missing sessionId field")
		}

		for _, s := range mockGameSessions {
			if s.ID == gameSessionId {
				return &Response{
					Status:  Success,
					Event:   EventGetGameSession,
					Payload: s,
				}, nil
			}
		}

		return nil, fmt.Errorf("game session not found")

	case string(EventNext):
		gameSessionID, found := message.Payload.(map[string]interface{})["sessionId"].(string)
		if !found {
			return nil, fmt.Errorf("missing sessionId field")
		}

		var gameSession GameSession

		for i, s := range mockGameSessions {
			if s.ID == gameSessionID {
				if mockGameSessions[i].Config.CurrentStage < len(mockGameSessions[i].Config.Stages)-1 {
					mockGameSessions[i].Config.CurrentStage++
				} else {
					mockGameSessions[i].Status = Finished
				}
				gameSession = mockGameSessions[i]
			}
		}

		return &Response{
			Status:  Success,
			Event:   EventGetGameSession,
			Payload: gameSession,
		}, nil

	case string(EventAnswer):
		jsonData, err := json.Marshal(message.Payload)
		if err != nil {
			return nil, err
		}

		var answer AnswerPayload
		err = json.Unmarshal(jsonData, &answer)
		if err != nil {
			return nil, err
		}

		for i, s := range mockGameSessions {
			if s.ID == answer.GameID {
				stage := mockGameSessions[i].Config.Stages[mockGameSessions[i].Config.CurrentStage]
				if stage.AnswerIndex == answer.OptionIndex {
					mockGameSessions[i].PlayerIDScoreMap[answer.PlayerID] += stage.Points
				}
			}
		}

		return &Response{
			Status:  Success,
			Event:   EventAnswer,
			Payload: nil,
		}, nil
	}

	return nil, nil
}

func getPermissionByEvent(event string) (Permission, error) {
	switch event {
	case string(EventCreateGameSession):
		return PermissionManageGameSession, nil
	case string(EventNext):
		return PermissionManageGameSession, nil
	case string(EventJoinGameQueue):
		return PermissionPlayGame, nil
	case string(EventLeaveGameQueue):
		return PermissionPlayGame, nil
	case string(EventGetGameSession):
		return PermissionGetGameSession, nil
	case string(EventAnswer):
		return PermissionPlayGame, nil
	default:
		return "", ErrorInvalidEvent
	}
}
