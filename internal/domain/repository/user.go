package repository

import (
	"database/sql"
	"myChat/internal/domain/model"
)

type UserRepositorier interface {
	Save(model.User) error
	FindById(int) (model.User, error)
	FindByEmail(string) (model.User, error)
	FindByUuid(string) (model.User, error)
	DeleteById(int) error
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

	return nil
}

func (ur *UserRepository) FindById(id int) (model.User, error) {
	var user model.User

	err := ur.db.QueryRow("SELECT * FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User

	err := ur.db.QueryRow("SELECT * FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) FindByUuid(uuid string) (model.User, error) {
	var user model.User

	err := ur.db.QueryRow("SELECT * FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Delete user from database
func (ur *UserRepository) DeleteById(id int) error {

	q := "DELETE FROM users WHERE id = $1"
	stmt, err := ur.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
