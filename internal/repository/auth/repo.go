package authrepo

import (
	repo "github.com/wrtgvr/microblog/internal/repository"
	"gorm.io/gorm"
)

type AuthRepo interface {
	TestFn() string
}

type PostgresAuthRepo struct {
	db *gorm.DB
}

func NewPostgresAuthRepo(db *repo.Database) *PostgresAuthRepo {
	return &PostgresAuthRepo{
		db: db.Conn,
	}
}

func (r *PostgresAuthRepo) TestFn() string {
	return "Hi there 3"
}
