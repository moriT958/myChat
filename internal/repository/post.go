package repository

import (
	"database/sql"
	"myChat/internal/model"
)

type PostRepositorier interface {
	Save(model.Post) error
	FindById(int) (model.Post, error)
	FindByThreadId(int) ([]model.Post, error)
	DeleteById(int) error
	DeleteByThreadId(int) error
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (pr *PostRepository) Save(post model.Post) error {
	tx, err := pr.db.Begin()
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

	var exists bool
	if err := tx.QueryRow("SELECT EXITS(SELECT 1 FROM posts WHERE id = $1)", post.Id).Scan(&exists); err != nil {
		return err
	}
	if exists {
		_, err = tx.Exec(
			"UPDATE posts SET uuid = $1, body = $2, user_id = $3, thread_id = $4, created_at = $5 WHERE id = $6",
			post.Uuid, post.Body, post.UserId, post.ThreadId, post.CreatedAt, post.Id,
		)
	} else {
		err = tx.QueryRow(
			"INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			post.Uuid, post.Body, post.UserId, post.ThreadId, post.CreatedAt,
		).Scan(&post.Id)
	}

	return nil
}

func (pr *PostRepository) FindById(id int) (model.Post, error) {
	var post model.Post
	err := pr.db.QueryRow("SELECT * FROM posts WHERE id = $1", id).Scan(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (pr *PostRepository) FindByThreadId(threadId int) ([]model.Post, error) {
	var posts []model.Post
	rows, err := pr.db.Query("SELECT * FROM posts WHERE thread_id = $1", threadId)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		post := model.Post{}
		if err := rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	rows.Close()

	return posts, nil
}

func (pr *PostRepository) DeleteById(id int) error {
	_, err := pr.db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRepository) DeleteByThreadId(threadId int) error {
	_, err := pr.db.Exec("DELETE FROM posts WHERE thread_id = $1", threadId)
	if err != nil {
		return err
	}
	return nil
}
