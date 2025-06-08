package userrepo

import "gorm.io/gorm"

type UserRepo interface {
	TestFn() string
}

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewPostgresUserRepo(db *gorm.DB) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

func (r *PostgresUserRepo) TestFn() string {
	return "Hi there"
}
