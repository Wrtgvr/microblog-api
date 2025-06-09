package userhandler

import (
	"github.com/wrtgvr/microblog/internal/handlers"
	repo "github.com/wrtgvr/microblog/internal/repository"
	userrepo "github.com/wrtgvr/microblog/internal/repository/user"
	userservice "github.com/wrtgvr/microblog/internal/services/user"
)

func NewUserHandlerWithDeps(db *repo.Database, deps *handlers.HandlerDeps) *UserHandler {
	repo := userrepo.NewPostgresUserRepo(db)
	service := userservice.NewUserService(repo)

	return NewUserHandler(deps, service)
}
