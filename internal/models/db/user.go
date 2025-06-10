package modelsdb

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64 `gorm:"primaryKey"`
	Username     string `gorm:"size:30;unique;not null"`
	PasswordHash string `gorm:"size:60;not null"`
	Posts        []Post `gorm:"foreignKey:UserID;constraint;OnDelete:CASCADE"`
	Role         string `gorm:"size:16;not null;default:'user'"` // user, moderator, admin
	IsActive     bool   `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
