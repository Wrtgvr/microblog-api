package app

import (
	"fmt"
	"os"

	"github.com/wrtgvr/microblog/internal/handlers"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
	repo "github.com/wrtgvr/microblog/internal/repository"
	"github.com/wrtgvr/microblog/internal/router"
	"github.com/wrtgvr/microblog/internal/server"
)

type App struct {
	Server *server.Server
}

func NewApp(port int) *App {
	errHandler := prepareGinErrorHandler()
	handlerDeps := handlers.NewHandlerDeps(errHandler)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)

	db := repo.MustOpenDb(dsn)

	userHandler := userhandler.NewUserHandlerWithDeps(db, handlerDeps)

	r := router.NewRouter()
	router.RegisterUserRoutes(r, userHandler)

	srv := server.NewServer(port, r)

	app := &App{
		Server: srv,
	}

	return app
}
