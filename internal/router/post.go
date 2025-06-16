package router

import (
	"github.com/gin-gonic/gin"
	posthandler "github.com/wrtgvr/microblog/internal/handlers/post"
)

func RegisterPostsRoutes(r *gin.Engine, h *posthandler.PostHandler) {
	g := r.Group("/posts")
	{
		g.GET("/:id", func(ctx *gin.Context) {})            // get post with given id -GUEST ROLE
		g.DELETE("/:id", authMw, func(ctx *gin.Context) {}) // delete post with given id -ADMIN ROLE (if user is post author then -USER ROLE)

		g.POST("", authMw, func(ctx *gin.Context) {}) // create new post -USER ROLE
	}
}
