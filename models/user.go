package models

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
func (m *Models) CreateSession(u User) (session Session, err error) {
	q := "INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, email, user_id, created_at"
	stmt, err := m.db.Prepare(q)
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
func (m *Models) GetSessionByUser(u User) (session Session, err error) {
	session = Session{}
	err = m.db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", u.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Create a new user, save user info into the database
func (m *Models) CreateUser(u User) (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	q := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"
	stmt, err := m.db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now()).Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Delete user from database
func (m *Models) DeleteUser(u User) (err error) {
	q := "DELETE FROM users WHERE id = $1"
	stmt, err := m.db.Prepare(q)
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
func (m *Models) UpdateUser(u User) (err error) {
	q := "UPDATE users SET name = $2, email = $3 WHERE id = $1"
	stmt, err := m.db.Prepare(q)
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

// Delete all users from database
func (m *Models) DeleteAllUsers(u User) (err error) {
	q := "DELETE FROM users"
	_, err = m.db.Exec(q)
	if err != nil {
		return
	}
	return
}

// Get all users in the database and returns it
func (m *Models) GetAllUsers() (users []User, err error) {
	rows, err := m.db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// Get a single user given the email
func (m *Models) GetUserByEmail(u User, email string) (user User, err error) {
	user = User{}
	err = m.db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Get a single user given the UUID
func (m *Models) GetUserByUUID(u User, uuid string) (user User, err error) {
	user = User{}
	err = m.db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Create a new thread
func (m *Models) CreateThread(u User, topic string) (conv Thread, err error) {
	q := "INSERT INTO threads (uuid, topic, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, topic, user_id, created_at"
	stmt, err := m.db.Prepare(q)
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
func (m *Models) CreatePost(u User, conv Thread, body string) (post Post, err error) {
	q := "INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, body, user_id, thread_id, created_at"
	stmt, err := m.db.Prepare(q)
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
