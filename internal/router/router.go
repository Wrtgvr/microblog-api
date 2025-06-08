package router

import (
	"github.com/gin-gonic/gin"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	return r
}

func RegisterUserRoutes(r *gin.Engine, h *userhandler.UserHandler) {
	//g := r.Group("users")
}
