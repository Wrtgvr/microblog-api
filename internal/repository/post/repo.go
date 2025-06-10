package postrepo

import (
	repo "github.com/wrtgvr/microblog/internal/repository"
	"gorm.io/gorm"
)

type PostRepo interface {
	TestFn() string
}

type PostgresPostRepo struct {
	db *gorm.DB
}

func NewPostgresPostRepo(db *repo.Database) *PostgresPostRepo {
	return &PostgresPostRepo{
		db: db.Conn,
	}
}

func (r *PostgresPostRepo) TestFn() string {
	return "Hi there 2"
}
