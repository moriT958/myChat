package repository

import (
	"time"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// Check if session is valid in the database
func (s *Session) Check() (res bool, err error) {
	db = GetDB()
	err = db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", s.Uuid).
		Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	if err != nil {
		return
	}
	if s.Id != 0 {
		return
	}
	return
}

// Delete session from database
func (s *Session) Delete() (err error) {
	q := "DELETE FROM sessions WHERE uuid = $1"
	db = GetDB()
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.Uuid)
	if err != nil {
		return
	}
	return
}

// Get the user from the session
func (s *Session) GetUser() (usr User, err error) {
	db = GetDB()
	err = db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", s.UserId).
		Scan(&usr.Id, &usr.Uuid, &usr.Name, &usr.Email, &usr.CreatedAt)
	if err != nil {
		return
	}
	return
}
