package router

import (
	"github.com/gin-gonic/gin"
	authhandler "github.com/wrtgvr/microblog/internal/handlers/auth"
)

func RegisterAuthRoutes(r *gin.Engine, h *authhandler.AuthHandler) {
	g := r.Group("/auth")
	{
		g.POST("/register", resfreshAuthMw, func(ctx *gin.Context) {})
		g.POST("/login", resfreshAuthMw, func(ctx *gin.Context) {})
	}
}
