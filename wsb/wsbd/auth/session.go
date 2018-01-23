package auth

type contextKey string

type Session struct {
	Identifier string
}

const SessionContextKey contextKey = "Session"
