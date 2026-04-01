package main

import (
	"net/http"
	"strings"

	"github.com/buemura/gox/examples/social_media/views"
	"golang.org/x/crypto/bcrypt"
)

func handleLoginPage(w http.ResponseWriter, r *http.Request) {
	if getCurrentUser(r) != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.LoginPage("").Render(r.Context(), w)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.FormValue("username"))
	password := r.FormValue("password")

	user, hash, err := getUserByUsername(username)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.LoginPage("Invalid username or password.").Render(r.Context(), w)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.LoginPage("Invalid username or password.").Render(r.Context(), w)
		return
	}

	token, err := createSession(user.ID)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.LoginPage("Something went wrong. Please try again.").Render(r.Context(), w)
		return
	}

	setSessionCookie(w, token)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleRegisterPage(w http.ResponseWriter, r *http.Request) {
	if getCurrentUser(r) != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.RegisterPage("").Render(r.Context(), w)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	displayName := strings.TrimSpace(r.FormValue("display_name"))
	username := strings.TrimSpace(r.FormValue("username"))
	email := strings.TrimSpace(r.FormValue("email"))
	password := r.FormValue("password")

	if displayName == "" || username == "" || email == "" || password == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.RegisterPage("All fields are required.").Render(r.Context(), w)
		return
	}

	if len(password) < 6 {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.RegisterPage("Password must be at least 6 characters.").Render(r.Context(), w)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.RegisterPage("Something went wrong. Please try again.").Render(r.Context(), w)
		return
	}

	userID, err := createUser(username, displayName, email, string(hash))
	if err != nil {
		msg := "Something went wrong. Please try again."
		if strings.Contains(err.Error(), "UNIQUE") {
			msg = "Username or email already taken."
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.RegisterPage(msg).Render(r.Context(), w)
		return
	}

	token, err := createSession(int(userID))
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		views.LoginPage("Account created! Please sign in.").Render(r.Context(), w)
		return
	}

	setSessionCookie(w, token)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("session_token"); err == nil {
		deleteSession(cookie.Value)
	}
	clearSessionCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
