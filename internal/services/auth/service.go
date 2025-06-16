package authservice

import authrepo "github.com/wrtgvr/microblog/internal/repository/auth"

type AuthService struct {
	repo *authrepo.AuthRepo
}

func NewAuthService(repo authrepo.AuthRepo) *AuthService {
	return &AuthService{
		repo: &repo,
	}
}
