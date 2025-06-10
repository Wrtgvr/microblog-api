package app

import (
	"github.com/wrtgvr/microblog/internal/config"
	repo "github.com/wrtgvr/microblog/internal/repository"
	"github.com/wrtgvr/microblog/internal/server"
)

type App struct {
	Server *server.Server
}

func NewApp(cfg *config.Config) *App {
	db := repo.MustOpenDb(cfg.DB.GetDSN())

	userHandler, postHandler := initHandlers(db, prepareGinErrorHandler())

	r := setupRouter(userHandler, postHandler)

	return &App{
		Server: server.NewServer(cfg.Server.Port, r),
	}
}
