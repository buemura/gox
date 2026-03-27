package main

import (
	"context"
	"net/http"

	"github.com/buemura/gox/examples/social_media/views"
)

type contextKey string

const userContextKey contextKey = "user"

func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getCurrentUser(r)
		if user == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), userContextKey, user)
		next(w, r.WithContext(ctx))
	}
}

func getCurrentUser(r *http.Request) *views.User {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil
	}
	userID, err := getUserIDBySession(cookie.Value)
	if err != nil {
		return nil
	}
	user, err := getUserByID(userID)
	if err != nil {
		return nil
	}
	return user
}

func userFromContext(r *http.Request) *views.User {
	u, _ := r.Context().Value(userContextKey).(*views.User)
	return u
}

func setSessionCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400 * 30,
	})
}

func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
}
