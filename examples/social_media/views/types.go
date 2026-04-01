package views

import (
	"crypto/md5"
	"fmt"
	"strings"
)

type User struct {
	ID             int
	Username       string
	DisplayName    string
	Email          string
	Bio            string
	FollowerCount  int
	FollowingCount int
	PostCount      int
}

type PostData struct {
	ID           int
	UserID       int
	Username     string
	DisplayName  string
	Content      string
	TimeAgo      string
	ReplyToID    int
	ReplyToUser  string
	LikeCount    int
	CommentCount int
	ReplyCount   int
	Liked        bool
}

type CommentData struct {
	ID          int
	Username    string
	DisplayName string
	Content     string
	TimeAgo     string
}

type FeedPageData struct {
	Posts []PostData
}

type PostDetailData struct {
	Post     PostData
	Thread   []PostData
	Replies  []PostData
	Comments []CommentData
}

type ProfileData struct {
	ProfileUser  User
	IsFollowing  bool
	IsOwnProfile bool
	Posts        []PostData
}

type ExploreData struct {
	Posts          []PostData
	SuggestedUsers []User
}

func AvatarColor(username string) string {
	colors := []string{
		"#1d9bf0", "#00ba7c", "#f91880", "#ffd400",
		"#7856ff", "#ff7a00", "#00d5fa", "#794bc4",
	}
	hash := md5.Sum([]byte(username))
	idx := int(hash[0]) % len(colors)
	return colors[idx]
}

func AvatarInitial(username string) string {
	if len(username) == 0 {
		return "?"
	}
	return strings.ToUpper(username[:1])
}

func FormatCount(n int) string {
	if n >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(n)/1000000)
	}
	if n >= 1000 {
		return fmt.Sprintf("%.1fK", float64(n)/1000)
	}
	return fmt.Sprintf("%d", n)
}

func IsReply(post PostData) bool {
	return post.ReplyToID > 0
}
