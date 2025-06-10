package app

import (
	"github.com/gin-gonic/gin"
	posthandler "github.com/wrtgvr/microblog/internal/handlers/post"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
	"github.com/wrtgvr/microblog/internal/router"
)

func setupRouter(userHandler *userhandler.UserHandler, postHandler *posthandler.PostHandler) *gin.Engine {
	r := router.NewRouter()

	router.RegisterUserRoutes(r, userHandler)
	router.RegisterPostsRoutes(r, postHandler)

	return r
}
