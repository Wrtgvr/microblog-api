package router

import (
	"github.com/gin-gonic/gin"
	posthandler "github.com/wrtgvr/microblog/internal/handlers/post"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	return r
}

func RegisterUserRoutes(r *gin.Engine, h *userhandler.UserHandler) {
	g := r.Group("users")
	{
		g.GET("/me", func(ctx *gin.Context) {})       // user get their account info -USER ROLE
		g.GET("/me/posts", func(ctx *gin.Context) {}) // get user's posts -USER ROLE
		g.PATCH("/me", func(ctx *gin.Context) {})     // user updates their account info -USER ROLE
		g.DELETE("/me", func(ctx *gin.Context) {})    // user delete their accout -USER ROLE

		g.GET("/:id", func(ctx *gin.Context) {})       // get user account info with given id -GUEST ROLE
		g.GET("/:id/posts", func(ctx *gin.Context) {}) // get user's posts -GUEST ROLE
		g.PATCH("/:id", func(ctx *gin.Context) {})     // update account info with given id -ADMIN ROLE (if user is acc owner then -USER ROLE)
		g.DELETE("/:id", func(ctx *gin.Context) {})    // delete account info with given id -ADMIN ROLE (if user is acc owner then -USER ROLE)
	}
}

func RegisterPostsRoutes(r *gin.Engine, h *posthandler.PostHandler) {
	g := r.Group("posts")
	{
		g.GET("/:id", func(ctx *gin.Context) {})    // get post with given id -GUEST ROLE
		g.DELETE("/:id", func(ctx *gin.Context) {}) // delete post with given id -ADMIN ROLE (if user is post author then -USER ROLE)

		g.POST("", func(ctx *gin.Context) {}) // create new post -USER ROLE
	}
}
