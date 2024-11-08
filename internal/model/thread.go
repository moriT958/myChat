package model

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
	Posts     []Post
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// format the CreatedAt date to display nicely on the screen
func (t *Thread) CreatedAtStr() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (p *Post) CreatedAtStr() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}
