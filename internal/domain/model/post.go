package model

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

func (p *Post) CreatedAtStr() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}
