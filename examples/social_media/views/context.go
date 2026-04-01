package views

import "context"

type contextKey string

// UserContextKey is the context key used to store the current user.
// Set this in HTTP middleware via context.WithValue.
const UserContextKey contextKey = "user"

// CurrentUser extracts the authenticated user from the context.
// Returns a zero-value User if no user is set (e.g., unauthenticated pages).
func CurrentUser(ctx context.Context) User {
	u, _ := ctx.Value(UserContextKey).(*User)
	if u == nil {
		return User{}
	}
	return *u
}
