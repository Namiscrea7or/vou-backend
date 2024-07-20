package quiz

var mockIDTokenUserIDMap = map[string]string{
	"user":  "user",
	"admin": "admin",
}

var mockUserIDRoleMap = map[string]Role{
	"user":  User,
	"admin": Admin,
}

var mockGameSessions = []GameSession{}
