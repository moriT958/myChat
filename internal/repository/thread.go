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

	// save posts
	for _, post := range thread.Posts {
		// check if post exits.
		var postExists bool
		err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE id = $1)", post.Id).Scan(&postExists)
		if err != nil {
			return err
		}

		if postExists {
			// update post.
			_, err = tx.Exec(
				"UPDATE posts SET uuid = $1, body = $2, user_id = $3, thread_id = $4, created_at = $5 WHERE id = $6",
				post.Uuid, post.Body, post.UserId, post.ThreadId, post.CreatedAt, post.Id,
			)
		} else {
			// insert new post
			err = tx.QueryRow(
				"INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
				post.Uuid, post.Body, post.UserId, post.ThreadId, post.CreatedAt,
			).Scan(&post.Id)
		}

		if err != nil {
			return err
		}
	}

	// commit transaction
	return tx.Commit()
}

func (tr *ThreadRepository) FindById(id int) (model.Thread, error) {
	tx, err := tr.db.Begin()
	if err != nil {
		return model.Thread{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var thread model.Thread
	err = tx.QueryRow("SELECT * FROM threads WHERE id = $1", id).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	if err != nil {
		return model.Thread{}, err
	}

	// get posts replied on this thread
	posts, err := tr.getPosts(tx, id)
	if err != nil {
		return model.Thread{}, err
	}
	thread.Posts = posts

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
func (tr *ThreadRepository) CountPostNum(threadId int) (num int, err error) {
	rows, err := tr.db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", threadId)
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

func (tr *ThreadRepository) getPosts(tx *sql.Tx, threadId int) ([]model.Post, error) {
	rows, err := tx.Query("SELECT * FROM posts WHERE thread_id = $1", threadId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		post := model.Post{}
		if err := rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
