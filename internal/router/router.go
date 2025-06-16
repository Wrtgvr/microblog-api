package router

import (
	"github.com/gin-gonic/gin"
	authhandler "github.com/wrtgvr/microblog/internal/handlers/auth"
	posthandler "github.com/wrtgvr/microblog/internal/handlers/post"
	userhandler "github.com/wrtgvr/microblog/internal/handlers/user"
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

func RegisterPostsRoutes(r *gin.Engine, h *posthandler.PostHandler) {
	g := r.Group("/posts")
	{
		g.GET("/:id", func(ctx *gin.Context) {})            // get post with given id -GUEST ROLE
		g.DELETE("/:id", authMw, func(ctx *gin.Context) {}) // delete post with given id -ADMIN ROLE (if user is post author then -USER ROLE)

		g.POST("", authMw, func(ctx *gin.Context) {}) // create new post -USER ROLE
	}
}

func RegisterAuthRoutes(r *gin.Engine, h *authhandler.AuthHandler) {
	g := r.Group("/auth")
	{
		g.POST("/register", resfreshAuthMw, func(ctx *gin.Context) {})
		g.POST("/login", resfreshAuthMw, func(ctx *gin.Context) {})
	}
}
