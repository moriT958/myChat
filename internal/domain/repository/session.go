package repository

import (
	"database/sql"
	"myChat/internal/domain/model"
)

type SessionRepositorier interface {
	Save(model.Session) error
	FindByUuid(string) (model.Session, error)
	FindByUserId(int) ([]model.Session, error)
	DeleteByUuid(string) error
	DeleteByUserId(int) error
}

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (sr *SessionRepository) Save(session model.Session) error {
	// start db transaction
	tx, err := sr.db.Begin()
	if err != nil {
		return err
	}
	// rollback if error occurs.
	// and commit if proc finish correctly.
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// check if session exits in database
	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE id = $1)", session.Id).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		// update session.
		_, err = tx.Exec(
			"UPDATE sessions SET uuid = $1, email = $2, user_id = $3, created_at = $4 WHERE id = $5",
			session.Uuid, session.Email, session.UserId, session.CreatedAt, session.Id,
		)
	} else {
		// insert new session.
		err = tx.QueryRow(
			"INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
			session.Uuid, session.Email, session.UserId, session.CreatedAt,
		).Scan(&session.Id)
	}
	if err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepository) FindByUuid(uuid string) (model.Session, error) {
	var session model.Session

	err := sr.db.QueryRow("SELECT * FROM sessions WHERE uuid = $1", uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return session, err
	}

	return session, nil
}

func (sr *SessionRepository) FindByUserId(userId int) ([]model.Session, error) {
	var sessions []model.Session
	rows, err := sr.db.Query("SELECT * FROM sessions WHERE user_id = $1", userId)
	if err != nil {
		return sessions, err
	}

	for rows.Next() {
		var conv model.Session
		if err := rows.Scan(&conv.Id, &conv.Uuid, &conv.Email, &conv.UserId, &conv.CreatedAt); err != nil {
			return sessions, err
		}
		sessions = append(sessions, conv)
	}
	rows.Close()

	return sessions, nil
}

func (sr *SessionRepository) DeleteByUuid(uuid string) error {
	if _, err := sr.db.Exec("DELETE FROM session WHERE uuid = $1", uuid); err != nil {
		return err
	}
	return nil
}

func (sr *SessionRepository) DeleteByUserId(userId int) error {
	if _, err := sr.db.Exec("DELETE FROM sessions WHERE user_id = $1", userId); err != nil {
		return err
	}
	return nil
}
