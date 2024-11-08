package repository

import (
	"database/sql"
	"myChat/internal/model"
)

type UserRepositorier interface {
	Save(model.User) error
	FindById(int) (model.User, error)
	FindByEmail(string) (model.User, error)
	FindByUuid(string) (model.User, error)
	DeleteById(int) error
	DeleteSessionById(int) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Save(user model.User) error {
	// start db transaction
	tx, err := ur.db.Begin()
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

	// check if user exits in database
	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", user.Id).Scan(&exists)
	if err != nil {
		return err
	}

	// if user existed, update user.
	// if not exited, insert new user.
	if exists {
		_, err = tx.Exec(
			"UPDATE users SET uuid = $1, name = $2, email = $3, password = $4, created_at = $5 WHERE id = $6",
			user.Uuid, user.Name, user.Email, user.Password, user.CreatedAt, user.Id,
		)
	} else {
		err = tx.QueryRow(
			"INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			user.Uuid, user.Name, user.Email, user.Password, user.CreatedAt,
		).Scan(&user.Id)
	}
	if err != nil {
		return err
	}

	// save sessions
	for _, session := range user.Sessions {
		// check if session exits.
		var sessionExists bool
		err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE id = $1)", session.Id).Scan(&sessionExists)
		if err != nil {
			return err
		}
		if sessionExists {
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
	}

	return nil
}

func (ur *UserRepository) FindById(id int) (model.User, error) {
	var user model.User

	tx, err := ur.db.Begin()
	if err != nil {
		return user, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.QueryRow("SELECT * FROM users WHERE id == $1", id).Scan(&user)
	if err != nil {
		return user, err
	}

	// set sessions to user, if sessions exit.
	sessions, err := ur.getSessions(tx, id)
	if err != nil {
		return user, err
	}
	user.Sessions = sessions

	return user, nil
}

func (ur *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	tx, err := ur.db.Begin()
	if err != nil {
		return user, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.QueryRow("SELECT * FROM users WHERE email == $1", email).Scan(&user)
	if err != nil {
		return user, err
	}

	// set sessions to user, if sessions exit.
	sessions, err := ur.getSessions(tx, user.Id)
	if err != nil {
		return user, err
	}
	user.Sessions = sessions

	return user, nil
}

func (ur *UserRepository) FindByUuid(uuid string) (model.User, error) {
	var user model.User
	tx, err := ur.db.Begin()
	if err != nil {
		return user, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.QueryRow("SELECT * FROM users WHERE uuid == $1", uuid).Scan(&user)
	if err != nil {
		return user, err
	}

	// set sessions to user, if sessions exit.
	sessions, err := ur.getSessions(tx, user.Id)
	if err != nil {
		return user, err
	}
	user.Sessions = sessions

	return user, nil
}

// Delete user from database
func (ur *UserRepository) DeleteById(id int) error {
	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	q := "DELETE FROM users WHERE id = $1"
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	// delete sessions that this user has.
	_, err = tx.Exec("DELETE FROM sessions WHERE user_id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) getSessions(tx *sql.Tx, userId int) ([]model.Session, error) {
	var sessions []model.Session
	rows, err := tx.Query("SELECT * FROM sessions WHERE user_id == $1", userId)
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

// func (ur *UserRepository) DeleteSessionById(userId int) error {

// }
