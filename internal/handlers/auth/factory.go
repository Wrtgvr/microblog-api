package authhandler

import (
	"github.com/wrtgvr/microblog/internal/handlers"
	repo "github.com/wrtgvr/microblog/internal/repository"
	authrepo "github.com/wrtgvr/microblog/internal/repository/auth"
	authservice "github.com/wrtgvr/microblog/internal/services/auth"
)

func NewAuthHandlerWithDeps(db *repo.Database, deps *handlers.HandlerDeps) *AuthHandler {
	repo := authrepo.NewPostgresAuthRepo(db)
	service := authservice.NewAuthService(repo)

	return NewAuthHandler(deps, service)
}
