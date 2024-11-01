package models

import (
	"time"
)

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// format the CreatedAt date to display nicely on the screen
func (p *Post) CreatedAtStr() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// Get the user who wrote the post
func (p *Post) GetUser(db DbDependency) User {
	user := User{}
	db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", p.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return user
}
