package userservice

import userrepo "github.com/wrtgvr/microblog/internal/repository/user"

type UserService struct {
	repo *userrepo.UserRepo
}

func NewUserService(repo userrepo.UserRepo) *UserService {
	return &UserService{
		repo: &repo,
	}
}
