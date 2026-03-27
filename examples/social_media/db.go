package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math"
	"time"

	"github.com/buemura/gox/examples/social_media/views"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite", "social_media.db?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)")
	if err != nil {
		return err
	}
	return migrate()
}

func migrate() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			display_name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			bio TEXT DEFAULT '',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS sessions (
			token TEXT PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			content TEXT NOT NULL,
			reply_to_id INTEGER REFERENCES posts(id) ON DELETE SET NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS follows (
			follower_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			following_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (follower_id, following_id)
		);

		CREATE TABLE IF NOT EXISTS likes (
			user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (user_id, post_id)
		);

		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return err
}

// --- User queries ---

func createUser(username, displayName, email, passwordHash string) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO users (username, display_name, email, password_hash) VALUES (?, ?, ?, ?)",
		username, displayName, email, passwordHash,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func getUserByUsername(username string) (*views.User, string, error) {
	var u views.User
	var hash string
	err := db.QueryRow(
		"SELECT id, username, display_name, email, bio, password_hash FROM users WHERE username = ?",
		username,
	).Scan(&u.ID, &u.Username, &u.DisplayName, &u.Email, &u.Bio, &hash)
	if err != nil {
		return nil, "", err
	}
	return &u, hash, nil
}

func getUserByID(id int) (*views.User, error) {
	var u views.User
	err := db.QueryRow(
		"SELECT id, username, display_name, email, bio FROM users WHERE id = ?",
		id,
	).Scan(&u.ID, &u.Username, &u.DisplayName, &u.Email, &u.Bio)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func getUserProfile(username string) (*views.User, error) {
	var u views.User
	err := db.QueryRow(`
		SELECT u.id, u.username, u.display_name, u.email, u.bio,
			(SELECT COUNT(*) FROM follows WHERE following_id = u.id) as follower_count,
			(SELECT COUNT(*) FROM follows WHERE follower_id = u.id) as following_count,
			(SELECT COUNT(*) FROM posts WHERE user_id = u.id) as post_count
		FROM users u WHERE u.username = ?
	`, username).Scan(&u.ID, &u.Username, &u.DisplayName, &u.Email, &u.Bio,
		&u.FollowerCount, &u.FollowingCount, &u.PostCount)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// --- Session queries ---

func createSession(userID int) (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	token := hex.EncodeToString(b)
	_, err := db.Exec("INSERT INTO sessions (token, user_id) VALUES (?, ?)", token, userID)
	return token, err
}

func getUserIDBySession(token string) (int, error) {
	var userID int
	err := db.QueryRow("SELECT user_id FROM sessions WHERE token = ?", token).Scan(&userID)
	return userID, err
}

func deleteSession(token string) error {
	_, err := db.Exec("DELETE FROM sessions WHERE token = ?", token)
	return err
}

// --- Post queries ---

func createPost(userID int, content string, replyToID int) (int64, error) {
	var res sql.Result
	var err error
	if replyToID > 0 {
		res, err = db.Exec(
			"INSERT INTO posts (user_id, content, reply_to_id) VALUES (?, ?, ?)",
			userID, content, replyToID,
		)
	} else {
		res, err = db.Exec(
			"INSERT INTO posts (user_id, content) VALUES (?, ?)",
			userID, content,
		)
	}
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func deletePost(postID, userID int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ? AND user_id = ?", postID, userID)
	return err
}

func getFeedPosts(userID int) ([]views.PostData, error) {
	rows, err := db.Query(`
		SELECT p.id, p.user_id, u.username, u.display_name, p.content,
			p.reply_to_id, COALESCE(ru.username, ''),
			p.created_at,
			(SELECT COUNT(*) FROM likes WHERE post_id = p.id) as like_count,
			(SELECT COUNT(*) FROM comments WHERE post_id = p.id) as comment_count,
			(SELECT COUNT(*) FROM posts WHERE reply_to_id = p.id) as reply_count,
			EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = ?) as liked
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN posts rp ON p.reply_to_id = rp.id
		LEFT JOIN users ru ON rp.user_id = ru.id
		WHERE p.user_id = ? OR p.user_id IN (SELECT following_id FROM follows WHERE follower_id = ?)
		ORDER BY p.created_at DESC
		LIMIT 50
	`, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func getExplorePosts(userID int) ([]views.PostData, error) {
	rows, err := db.Query(`
		SELECT p.id, p.user_id, u.username, u.display_name, p.content,
			p.reply_to_id, COALESCE(ru.username, ''),
			p.created_at,
			(SELECT COUNT(*) FROM likes WHERE post_id = p.id) as like_count,
			(SELECT COUNT(*) FROM comments WHERE post_id = p.id) as comment_count,
			(SELECT COUNT(*) FROM posts WHERE reply_to_id = p.id) as reply_count,
			EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = ?) as liked
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN posts rp ON p.reply_to_id = rp.id
		LEFT JOIN users ru ON rp.user_id = ru.id
		ORDER BY p.created_at DESC
		LIMIT 50
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func getUserPosts(profileUserID, currentUserID int) ([]views.PostData, error) {
	rows, err := db.Query(`
		SELECT p.id, p.user_id, u.username, u.display_name, p.content,
			p.reply_to_id, COALESCE(ru.username, ''),
			p.created_at,
			(SELECT COUNT(*) FROM likes WHERE post_id = p.id) as like_count,
			(SELECT COUNT(*) FROM comments WHERE post_id = p.id) as comment_count,
			(SELECT COUNT(*) FROM posts WHERE reply_to_id = p.id) as reply_count,
			EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = ?) as liked
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN posts rp ON p.reply_to_id = rp.id
		LEFT JOIN users ru ON rp.user_id = ru.id
		WHERE p.user_id = ?
		ORDER BY p.created_at DESC
		LIMIT 50
	`, currentUserID, profileUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func getPostByID(postID, currentUserID int) (*views.PostData, error) {
	var p views.PostData
	var createdAt string
	var replyToID sql.NullInt64
	var replyToUser string
	err := db.QueryRow(`
		SELECT p.id, p.user_id, u.username, u.display_name, p.content,
			p.reply_to_id, COALESCE(ru.username, ''),
			p.created_at,
			(SELECT COUNT(*) FROM likes WHERE post_id = p.id) as like_count,
			(SELECT COUNT(*) FROM comments WHERE post_id = p.id) as comment_count,
			(SELECT COUNT(*) FROM posts WHERE reply_to_id = p.id) as reply_count,
			EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = ?) as liked
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN posts rp ON p.reply_to_id = rp.id
		LEFT JOIN users ru ON rp.user_id = ru.id
		WHERE p.id = ?
	`, currentUserID, postID).Scan(
		&p.ID, &p.UserID, &p.Username, &p.DisplayName, &p.Content,
		&replyToID, &replyToUser,
		&createdAt,
		&p.LikeCount, &p.CommentCount, &p.ReplyCount, &p.Liked,
	)
	if err != nil {
		return nil, err
	}
	if replyToID.Valid {
		p.ReplyToID = int(replyToID.Int64)
		p.ReplyToUser = replyToUser
	}
	p.TimeAgo = timeAgo(createdAt)
	return &p, nil
}

func getPostReplies(postID, currentUserID int) ([]views.PostData, error) {
	rows, err := db.Query(`
		SELECT p.id, p.user_id, u.username, u.display_name, p.content,
			p.reply_to_id, COALESCE(ru.username, ''),
			p.created_at,
			(SELECT COUNT(*) FROM likes WHERE post_id = p.id) as like_count,
			(SELECT COUNT(*) FROM comments WHERE post_id = p.id) as comment_count,
			(SELECT COUNT(*) FROM posts WHERE reply_to_id = p.id) as reply_count,
			EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = ?) as liked
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN posts rp ON p.reply_to_id = rp.id
		LEFT JOIN users ru ON rp.user_id = ru.id
		WHERE p.reply_to_id = ?
		ORDER BY p.created_at ASC
	`, currentUserID, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanPosts(rows)
}

func getThreadChain(postID, currentUserID int) ([]views.PostData, error) {
	var chain []views.PostData
	currentID := postID
	for {
		post, err := getPostByID(currentID, currentUserID)
		if err != nil || post.ReplyToID == 0 {
			break
		}
		currentID = post.ReplyToID
		parent, err := getPostByID(currentID, currentUserID)
		if err != nil {
			break
		}
		chain = append([]views.PostData{*parent}, chain...)
	}
	return chain, nil
}

func getPostComments(postID int) ([]views.CommentData, error) {
	rows, err := db.Query(`
		SELECT c.id, u.username, u.display_name, c.content, c.created_at
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.created_at ASC
	`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []views.CommentData
	for rows.Next() {
		var c views.CommentData
		var createdAt string
		if err := rows.Scan(&c.ID, &c.Username, &c.DisplayName, &c.Content, &createdAt); err != nil {
			return nil, err
		}
		c.TimeAgo = timeAgo(createdAt)
		comments = append(comments, c)
	}
	return comments, rows.Err()
}

// --- Like queries ---

func toggleLike(userID, postID int) error {
	var exists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = ? AND post_id = ?)", userID, postID).Scan(&exists)
	if exists {
		_, err := db.Exec("DELETE FROM likes WHERE user_id = ? AND post_id = ?", userID, postID)
		return err
	}
	_, err := db.Exec("INSERT INTO likes (user_id, post_id) VALUES (?, ?)", userID, postID)
	return err
}

// --- Comment queries ---

func addComment(userID, postID int, content string) error {
	_, err := db.Exec("INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)", userID, postID, content)
	return err
}

// --- Follow queries ---

func toggleFollow(followerID, followingID int) error {
	var exists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = ? AND following_id = ?)",
		followerID, followingID).Scan(&exists)
	if exists {
		_, err := db.Exec("DELETE FROM follows WHERE follower_id = ? AND following_id = ?", followerID, followingID)
		return err
	}
	_, err := db.Exec("INSERT INTO follows (follower_id, following_id) VALUES (?, ?)", followerID, followingID)
	return err
}

func isFollowing(followerID, followingID int) bool {
	var exists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = ? AND following_id = ?)",
		followerID, followingID).Scan(&exists)
	return exists
}

func getSuggestedUsers(currentUserID int) ([]views.User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.username, u.display_name, u.bio,
			(SELECT COUNT(*) FROM follows WHERE following_id = u.id) as follower_count
		FROM users u
		WHERE u.id != ? AND u.id NOT IN (SELECT following_id FROM follows WHERE follower_id = ?)
		ORDER BY follower_count DESC
		LIMIT 5
	`, currentUserID, currentUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []views.User
	for rows.Next() {
		var u views.User
		if err := rows.Scan(&u.ID, &u.Username, &u.DisplayName, &u.Bio, &u.FollowerCount); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

// --- Helpers ---

func scanPosts(rows *sql.Rows) ([]views.PostData, error) {
	var posts []views.PostData
	for rows.Next() {
		var p views.PostData
		var createdAt string
		var replyToID sql.NullInt64
		var replyToUser string
		if err := rows.Scan(
			&p.ID, &p.UserID, &p.Username, &p.DisplayName, &p.Content,
			&replyToID, &replyToUser,
			&createdAt,
			&p.LikeCount, &p.CommentCount, &p.ReplyCount, &p.Liked,
		); err != nil {
			return nil, err
		}
		if replyToID.Valid {
			p.ReplyToID = int(replyToID.Int64)
			p.ReplyToUser = replyToUser
		}
		p.TimeAgo = timeAgo(createdAt)
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func timeAgo(createdAt string) string {
	t, err := time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return createdAt
	}
	d := time.Since(t)
	switch {
	case d.Seconds() < 60:
		return fmt.Sprintf("%ds", int(math.Max(1, d.Seconds())))
	case d.Minutes() < 60:
		return fmt.Sprintf("%dm", int(d.Minutes()))
	case d.Hours() < 24:
		return fmt.Sprintf("%dh", int(d.Hours()))
	case d.Hours() < 720:
		return fmt.Sprintf("%dd", int(d.Hours()/24))
	default:
		return t.Format("Jan 2")
	}
}
