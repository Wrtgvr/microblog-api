package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/microblog/internal/router"
)

func setupRouter(h *Handlers) *gin.Engine {
	r := router.NewRouter()

	router.RegisterUserRoutes(r, h.User)
	router.RegisterPostsRoutes(r, h.Post)
	router.RegisterAuthRoutes(r, h.Auth)

	return r
}
