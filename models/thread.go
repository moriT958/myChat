package models

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
func (t *Thread) NumReplies(m Models) (int, error) {
	var num int
	rows, err := m.db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", t.Id)
	if err != nil {
		return num, err
	}
	for rows.Next() {
		if err = rows.Scan(&num); err != nil {
			return num, err
		}
	}
	rows.Close()
	return num, nil
}

// get posts to a thread
func (t *Thread) GetPosts(m Models) ([]Post, error) {
	var posts []Post
	rows, err := m.db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_id = $1", t.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	rows.Close()
	return posts, nil
}

// Get the user who started this thread
func (t *Thread) GetUser(m Models) (user User) {
	user = User{}
	m.db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", t.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
