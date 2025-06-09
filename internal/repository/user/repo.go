package userrepo

import (
	repo "github.com/wrtgvr/microblog/internal/repository"
	"gorm.io/gorm"
)

type UserRepo interface {
	TestFn() string
}

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewPostgresUserRepo(db *repo.Database) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db.Conn,
	}
}

func (r *PostgresUserRepo) TestFn() string {
	return "Hi there"
}
