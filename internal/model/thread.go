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
}

// format the CreatedAt date to display nicely on the screen
func (t *Thread) CreatedAtStr() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}
