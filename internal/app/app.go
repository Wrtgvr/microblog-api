package app

import (
	"github.com/wrtgvr/microblog/internal/handlers"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
	"github.com/wrtgvr/microblog/internal/router"
	"github.com/wrtgvr/microblog/internal/server"
	"gorm.io/gorm"
)

type App struct {
	Server *server.Server
}

func NewApp(port int) *App {
	errHandler := prepareGinErrorHandler()
	handlerDeps := handlers.NewHandlerDeps(errHandler)

	db := gorm.DB{} // dummy db init, will be changed soon

	userHandler := userhandler.NewUserHandlerWithDeps(db, handlerDeps)

	r := router.NewRouter()
	router.RegisterUserRoutes(r, userHandler)

	srv := server.NewServer(port, r)

	app := &App{
		Server: srv,
	}

	return app
}
