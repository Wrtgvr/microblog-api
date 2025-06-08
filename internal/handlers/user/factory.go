package userhandler

import (
	"github.com/wrtgvr/microblog/internal/handlers"
	userrepo "github.com/wrtgvr/microblog/internal/repository/user"
	userservice "github.com/wrtgvr/microblog/internal/services/user"
	"gorm.io/gorm"
)

func NewUserHandlerWithDeps(db *gorm.DB, deps *handlers.HandlerDeps) *UserHandler {
	repo := userrepo.NewPostgresUserRepo(db)
	service := userservice.NewUserService(repo)

	return NewUserHandler(deps, service)
}
