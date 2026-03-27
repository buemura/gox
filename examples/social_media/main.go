package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := initDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	mux := http.NewServeMux()

	// Auth routes (no auth required)
	mux.HandleFunc("GET /login", handleLoginPage)
	mux.HandleFunc("POST /login", handleLogin)
	mux.HandleFunc("GET /register", handleRegisterPage)
	mux.HandleFunc("POST /register", handleRegister)
	mux.HandleFunc("POST /logout", handleLogout)

	// Feed & explore
	mux.HandleFunc("GET /{$}", requireAuth(handleFeed))
	mux.HandleFunc("GET /explore", requireAuth(handleExplore))

	// Posts
	mux.HandleFunc("POST /posts", requireAuth(handleCreatePost))
	mux.HandleFunc("GET /post/{id}", requireAuth(handlePostDetail))
	mux.HandleFunc("POST /post/{id}/delete", requireAuth(handleDeletePost))
	mux.HandleFunc("POST /post/{id}/like", requireAuth(handleLikePost))
	mux.HandleFunc("POST /post/{id}/comment", requireAuth(handleAddComment))
	mux.HandleFunc("POST /post/{id}/reply", requireAuth(handleReplyPost))

	// Social
	mux.HandleFunc("GET /user/{username}", requireAuth(handleProfile))
	mux.HandleFunc("POST /user/{username}/follow", requireAuth(handleFollow))

	fmt.Println("Server running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", mux))
}
