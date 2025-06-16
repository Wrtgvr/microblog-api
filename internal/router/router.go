package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/microblog/internal/middleware"
)

var authMw = middleware.AuthMiddleware()
var resfreshAuthMw = middleware.RefreshAuthMiddleware()

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	return r
}
