package auth

type ContextKey int

const (
	ProfileKey ContextKey = iota
	UserKey
	RegisterKey
	PostDataKey
)

type Profile struct {
	UID         string
	PhoneNumber string
}
