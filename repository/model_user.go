package repository

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Create a new session for an existing user
func (u *User) CreateSession() (session Session, err error) {
	db = GetDB()
	q := "INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, email, user_id, created_at"
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), u.Email, u.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Get the session for an existing user
func (u *User) GetSession() (session Session, err error) {
	session = Session{}
	db = GetDB()
	err = db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", u.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Create a new user, save user info into the database
func (u *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	db = GetDB()
	q := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now()).
		Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Delete user from database
func (u *User) Delete() (err error) {
	db = GetDB()
	q := "DELETE FROM users WHERE id = $1"
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id)
	if err != nil {
		return
	}
	return
}

// Update user information in the database
func (u *User) Update() (err error) {
	db = GetDB()
	q := "UPDATE users SET name = $2, email = $3 WHERE id = $1"
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id, u.Name, u.Email)
	if err != nil {
		return
	}
	return
}

// Create a new thread
func (u *User) CreateThread(topic string) (conv Thread, err error) {
	db = GetDB()
	q := "INSERT INTO threads (uuid, topic, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, topic, user_id, created_at"
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), topic, u.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Create a new post to a thread
func (u *User) CreatePost(conv Thread, body string) (post Post, err error) {
	db = GetDB()
	q := "INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, body, user_id, thread_id, created_at"
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), body, u.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	if err != nil {
		return
	}
	return
}
