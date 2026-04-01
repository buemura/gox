package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/buemura/gox/examples/social_media/views"
)

func handleFeed(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	posts, err := getFeedPosts(user.ID)
	if err != nil {
		http.Error(w, "Failed to load feed", http.StatusInternalServerError)
		return
	}
	if posts == nil {
		posts = []views.PostData{}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.FeedPage(views.FeedPageData{
		Posts: posts,
	}).Render(r.Context(), w)
}

func handleExplore(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	posts, _ := getExplorePosts(user.ID)
	if posts == nil {
		posts = []views.PostData{}
	}
	suggested, _ := getSuggestedUsers(user.ID)
	if suggested == nil {
		suggested = []views.User{}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.ExplorePage(views.ExploreData{
		Posts:          posts,
		SuggestedUsers: suggested,
	}).Render(r.Context(), w)
}

func handleCreatePost(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	content := strings.TrimSpace(r.FormValue("content"))
	if content == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	createPost(user.ID, content, 0)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handlePostDetail(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	post, err := getPostByID(postID, user.ID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	thread, _ := getThreadChain(postID, user.ID)
	if thread == nil {
		thread = []views.PostData{}
	}
	replies, _ := getPostReplies(postID, user.ID)
	if replies == nil {
		replies = []views.PostData{}
	}
	comments, _ := getPostComments(postID)
	if comments == nil {
		comments = []views.CommentData{}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.PostDetailPage(views.PostDetailData{
		Post:     *post,
		Thread:   thread,
		Replies:  replies,
		Comments: comments,
	}).Render(r.Context(), w)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	deletePost(postID, user.ID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleLikePost(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	toggleLike(user.ID, postID)
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = fmt.Sprintf("/post/%d", postID)
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}

func handleAddComment(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	content := strings.TrimSpace(r.FormValue("content"))
	if content != "" {
		addComment(user.ID, postID, content)
	}
	http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusSeeOther)
}

func handleReplyPost(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	content := strings.TrimSpace(r.FormValue("content"))
	if content == "" {
		http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusSeeOther)
		return
	}
	newID, err := createPost(user.ID, content, postID)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/post/%d", newID), http.StatusSeeOther)
}
