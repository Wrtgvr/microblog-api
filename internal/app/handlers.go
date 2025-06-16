package app

import (
	"github.com/wrtgvr/errsuit/drivers/ginadap"
	"github.com/wrtgvr/microblog/internal/handlers"
	authhandler "github.com/wrtgvr/microblog/internal/handlers/auth"
	posthandler "github.com/wrtgvr/microblog/internal/handlers/post"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
	repo "github.com/wrtgvr/microblog/internal/repository"
)

type Handlers struct {
	User *userhandler.UserHandler
	Post *posthandler.PostHandler
	Auth *authhandler.AuthHandler
}

func initHandlers(db *repo.Database, errHandler *ginadap.GinErrorHandler) *Handlers {
	handlerDeps := handlers.NewHandlerDeps()

	userHandler := userhandler.NewUserHandlerWithDeps(db, handlerDeps)
	postHandler := posthandler.NewPostHandlerWithDeps(db, handlerDeps)
	authHandler := authhandler.NewAuthHandlerWithDeps(db, handlerDeps)

	return &Handlers{
		userHandler, postHandler, authHandler,
	}
}
