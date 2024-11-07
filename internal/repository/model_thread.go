package repository

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// format the CreatedAt date to display nicely on the screen
func (t *Thread) CreatedAtStr() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (t *Thread) NumReplies() (num int, err error) {
	db := GetDB()
	rows, err := db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", t.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&num); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// get posts to a thread
func (t *Thread) GetPosts() (posts []Post, err error) {
	db = GetDB()
	rows, err := db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_id = $1", t.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get the user who started this thread
func (t *Thread) GetUser() (user User) {
	user = User{}
	db := GetDB()
	db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", t.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	return
}
