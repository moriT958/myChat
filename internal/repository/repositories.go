package repository

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Get and Delete functions below

// Get all threads in the database and returns it
func (r *Repository) GetAllThreads() (threads []Thread, err error) {
	rows, err := r.db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
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
func (r *Repository) GetThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = r.db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Delete all sessions from database
func (r *Repository) DeleteAllSessions() error {
	q := "DELETE FROM sessions"
	_, err := r.db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

// Delete all users from database
func (r *Repository) DeleteAllUsers() (err error) {
	q := "DELETE FROM users"
	_, err = r.db.Exec(q)
	if err != nil {
		return
	}
	return
}

// Get all users in the database and returns it
func (r *Repository) GetAllUsers() (users []User, err error) {
	rows, err := r.db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
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
func (r *Repository) GetUserByEmail(email string) (user User, err error) {
	user = User{}
	err = r.db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}

// Get a single user given the UUID
func (r *Repository) GetUserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = r.db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return
	}
	return
}
