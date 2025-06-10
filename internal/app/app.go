package app

import (
	"github.com/wrtgvr/microblog/internal/config"
	repo "github.com/wrtgvr/microblog/internal/repository"
	"github.com/wrtgvr/microblog/internal/server"
)

type App struct {
	Server *server.Server
}

func NewApp(port int) *App {
	dsn := config.GetDBConnectionString()
	db := repo.MustOpenDb(dsn)

	userHandler, postHandler := initHandlers(db, prepareGinErrorHandler())

	r := setupRouter(userHandler, postHandler)

	return &App{
		Server: server.NewServer(port, r),
	}
}
