package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/errsuit/drivers/ginadap"
	"github.com/wrtgvr/microblog/internal/router"
	"github.com/wrtgvr/microblog/internal/server"
)

type App struct {
	ErrHandler *ginadap.GinErrorHandler
	Server     *server.Server
}

func NewApp(port int) *App {
	errHandler := prepareGinErrorHandler()

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	srv := server.NewServer(port, r)
	router.RegisterRoutes(r)

	app := &App{
		ErrHandler: errHandler,
		Server:     srv,
	}

	return app
}
