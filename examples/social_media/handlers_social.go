package main

import (
	"fmt"
	"net/http"

	"github.com/buemura/gox/examples/social_media/views"
)

func handleFollow(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	username := r.PathValue("username")

	target, _, err := getUserByUsername(username)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if target.ID != user.ID {
		toggleFollow(user.ID, target.ID)
	}

	http.Redirect(w, r, fmt.Sprintf("/user/%s", username), http.StatusSeeOther)
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	username := r.PathValue("username")

	profileUser, err := getUserProfile(username)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	posts, _ := getUserPosts(profileUser.ID, user.ID)
	if posts == nil {
		posts = []views.PostData{}
	}

	following := isFollowing(user.ID, profileUser.ID)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.ProfilePage(views.ProfileData{
		ProfileUser:  *profileUser,
		IsFollowing:  following,
		IsOwnProfile: user.ID == profileUser.ID,
		Posts:        posts,
	}).Render(r.Context(), w)
}
