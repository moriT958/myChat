package models

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
func (s *Session) Check(db DbDependency) (bool, error) {
	err := db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", s.Uuid).
		Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	if err != nil {
		return false, err
	}
	if s.Id != 0 {
		return true, err
	}
	return true, nil
}

// Delete session from database
func (s *Session) DeleteByUUID(db DbDependency) error {
	q := "DELETE FROM sessions WHERE uuid = $1"
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.Uuid)
	if err != nil {
		return err
	}
	return nil
}

// Get the user from the session
func (s *Session) GetUser(db DbDependency) (User, error) {
	usr := User{}
	err := db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", s.UserId).
		Scan(&usr.Id, &usr.Uuid, &usr.Name, &usr.Email, &usr.CreatedAt)
	if err != nil {
		return usr, err
	}
	return usr, nil
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
