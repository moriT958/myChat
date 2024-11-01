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
func (t *Thread) NumReplies(db DbDependency) (int, error) {
	var num int
	rows, err := db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", t.Id)
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
func (t *Thread) GetPosts(db DbDependency) ([]Post, error) {
	var posts []Post
	rows, err := db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_id = $1", t.Id)
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

// Get all threads in the database and returns it
func (m *Models) GetAllThreads() (threads []Thread, err error) {
	rows, err := m.db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

// Get a thread by the UUID
func (m *Models) GetThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = m.db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Get the user who started this thread
func (t *Thread) GetUser(db DbDependency) (user User) {
	user = User{}
	db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", t.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
