package modelsdb

import (
	"time"
)

type Post struct {
	ID        uint64 `gorm:"primaryKey;index"`
	UserID    uint64 `gorm:"not null;index"`
	Title     string `gorm:"size:30"`
	Content   string `gorm:"size:300"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
