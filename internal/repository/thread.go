package repository

import (
	"database/sql"
	"myChat/internal/model"
)

type ThreadRepositorier interface {
	Save(model.Thread) error
	FindById(int) (model.Thread, error)
	FindAll() ([]model.Thread, error)
	DeleteById(int) error
	CountPostNum(int) (int, error)
}

type ThreadRepository struct {
	db *sql.DB
}

func NewThreadRepository(db *sql.DB) *ThreadRepository {
	return &ThreadRepository{db: db}
}

func (tr *ThreadRepository) Save(thread model.Thread) error {
	// start db transaction
	tx, err := tr.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		// rollback if error occurs.
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// check if thread exits in database
	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM threads WHERE id = $1)", thread.Id).Scan(&exists)
	if err != nil {
		return err
	}
	// if thread existed, update thread.
	// if not exited, insert new thread.
	if exists {
		_, err = tx.Exec(
			"UPDATE threads SET uuid = $1, topic = $2, user_id = $3, created_at = $4 WHERE id = $5",
			thread.Uuid, thread.Topic, thread.UserId, thread.CreatedAt, thread.Id,
		)
	} else {
		err = tx.QueryRow(
			"INSERT INTO threads (uuid, topic, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
			thread.Uuid, thread.Topic, thread.UserId, thread.CreatedAt,
		).Scan(&thread.Id)
	}
	if err != nil {
		return err
	}

	return nil
}

func (tr *ThreadRepository) FindById(id int) (model.Thread, error) {

	var thread model.Thread
	err := tr.db.QueryRow("SELECT * FROM threads WHERE id = $1", id).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	if err != nil {
		return model.Thread{}, err
	}

	return thread, nil
}

// Get all threads in the database and returns it
func (tr *ThreadRepository) FindAll() (threads []model.Thread, err error) {
	rows, err := tr.db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := model.Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

func (tr *ThreadRepository) DeleteById(id int) error {
	_, err := tr.db.Exec("DELETE FROM threads WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// get the number of posts in a thread
func (tr *ThreadRepository) CountPostNum(id int) (num int, err error) {
	rows, err := tr.db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&num); err != nil {
			return
		}
	}
	rows.Close()
	return
}
