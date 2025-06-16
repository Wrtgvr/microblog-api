package router

import (
	"github.com/gin-gonic/gin"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
)

func RegisterUserRoutes(r *gin.Engine, h *userhandler.UserHandler) {
	g := r.Group("/users")
	{
		meGroup := g.Group("/me")
		meGroup.Use(authMw)
		{
			g.GET("", func(ctx *gin.Context) {})       // user get their account info -USER ROLE
			g.GET("/posts", func(ctx *gin.Context) {}) // get user's posts -USER ROLE
			g.PATCH("", func(ctx *gin.Context) {})     // user updates their account info -USER ROLE
			g.DELETE("", func(ctx *gin.Context) {})    // user delete their accout -USER ROLE
		}
		g.GET("/:id", func(ctx *gin.Context) {})            // get user account info with given id -GUEST ROLE
		g.GET("/:id/posts", func(ctx *gin.Context) {})      // get user's posts -GUEST ROLE
		g.PATCH("/:id", authMw, func(ctx *gin.Context) {})  // update account info with given id -ADMIN ROLE (if user is acc owner then -USER ROLE)
		g.DELETE("/:id", authMw, func(ctx *gin.Context) {}) // delete account info with given id -ADMIN ROLE (if user is acc owner then -USER ROLE)
	}
}
