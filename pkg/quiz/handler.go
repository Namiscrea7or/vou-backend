package quiz

import (
	"context"

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
	}

	return nil, nil
}

func getPermissionByEvent(event string) (Permission, error) {
	switch event {
	case string(EventCreateGameSession):
		return PermissionManageGameSession, nil
	case string(EventUpdateGameSessionConfig):
		return PermissionManageGameSession, nil
	case string(EventJoinGameQueue):
		return PermissionPlayGame, nil
	case string(EventLeaveGameQueue):
		return PermissionPlayGame, nil
	default:
		return "", ErrorInvalidEvent
	}
}
