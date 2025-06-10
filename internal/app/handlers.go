package app

import (
	"github.com/wrtgvr/errsuit/drivers/ginadap"
	"github.com/wrtgvr/microblog/internal/handlers"
	posthandler "github.com/wrtgvr/microblog/internal/handlers/post"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
	repo "github.com/wrtgvr/microblog/internal/repository"
)

func initHandlers(db *repo.Database, errHandler *ginadap.GinErrorHandler) (
	*userhandler.UserHandler, *posthandler.PostHandler,
) {
	handlerDeps := handlers.NewHandlerDeps(errHandler)

	userHandler := userhandler.NewUserHandlerWithDeps(db, handlerDeps)
	postHandler := posthandler.NewPostHandlerWithDeps(db, handlerDeps)

	return userHandler, postHandler
}
