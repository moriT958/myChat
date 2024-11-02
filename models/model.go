package models

import (
	"database/sql"
)

type DbDependency interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
}

type Models struct {
	db DbDependency
}

func NewModels(db DbDependency) *Models {
	return &Models{db: db}
}

// Delete all sessions from database
func (m *Models) DeleteAllSessions() error {
	q := "DELETE FROM sessions"
	_, err := m.db.Exec(q)
	if err != nil {
		return err
	}
	return nil
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

// Delete all users from database
func (m *Models) DeleteAllUsers() (err error) {
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
func (m *Models) GetUserByEmail(email string) (user User, err error) {
	user = User{}
	err = m.db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Get a single user given the UUID
func (m *Models) GetUserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = m.db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}
