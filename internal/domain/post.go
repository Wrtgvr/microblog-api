package domain

import "time"

type Post struct {
	ID        uint64
	UserID    uint64
	Title     string
	Content   string
	CreatedAt time.Time
}
