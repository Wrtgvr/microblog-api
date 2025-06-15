package app

import (
	"github.com/wrtgvr/errsuit/drivers/ginadap"
	"github.com/wrtgvr/microblog/internal/config"
	repo "github.com/wrtgvr/microblog/internal/repository"
	"github.com/wrtgvr/microblog/internal/server"
)

type App struct {
	Server *server.Server
}

func NewApp(cfg *config.Config) *App {
	db := repo.MustOpenDb(cfg.DB.GetDSN())
	errHandler := prepareGinErrorHandler()

	userHandler, postHandler := initHandlers(db, errHandler)

	r := setupRouter(userHandler, postHandler)

	r.Use(ginadap.InjectErrHandlerMiddleware(errHandler))

	return &App{
		Server: server.NewServer(cfg.Server.Port, r),
	}
}
