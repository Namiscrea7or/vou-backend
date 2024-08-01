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
	userID, found := mockIDTokenUserIDMap[idToken]
	if !found {
		return nil, ErrorUserNotFound
	}

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

	case string(EventUpdateGameSession):
		jsonData, err := json.Marshal(message.Payload)
		if err != nil {
			return nil, err
		}

		var gameSession GameSession
		err = json.Unmarshal(jsonData, &gameSession)
		if err != nil {
			return nil, err
		}

		for i, s := range mockGameSessions {
			if s.ID == gameSession.ID {
				mockGameSessions[i] = gameSession
			}
		}

		return &Response{
			Status:  Success,
			Event:   EventGetGameSession,
			Payload: gameSession,
		}, nil

	}

	return nil, nil
}

func getPermissionByEvent(event string) (Permission, error) {
	switch event {
	case string(EventCreateGameSession):
		return PermissionManageGameSession, nil
	case string(EventUpdateGameSession):
		return PermissionManageGameSession, nil
	case string(EventJoinGameQueue):
		return PermissionPlayGame, nil
	case string(EventLeaveGameQueue):
		return PermissionPlayGame, nil
	case string(EventGetGameSession):
		return PermissionGetGameSession, nil
	default:
		return "", ErrorInvalidEvent
	}
}
