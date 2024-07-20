package quiz

func hasPermission(role Role, permission Permission) bool {
	permissions, ok := rolePermissionsMap[role]
	if !ok {
		return false
	}
	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}

var rolePermissionsMap = map[Role][]Permission{
	Admin: {
		PermissionManageGameSession,
	},
	User: {
		PermissionPlayGame,
	},
}
