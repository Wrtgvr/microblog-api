package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	errsuit "github.com/wrtgvr/errsuit/core"
	"github.com/wrtgvr/microblog/internal/pkg/jwt"
	"github.com/wrtgvr/microblog/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenPrefix := "Bearer "

		errHandleFunc := utils.MakeErrHandlerFunc(c)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, tokenPrefix) {
			errHandleFunc(errsuit.NewUnauthorized("missing authorization token", nil, false))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, tokenPrefix)
		err := jwt.ParseToken(tokenStr)
		if err != nil {
			errHandleFunc(errsuit.NewUnauthorized("invalid token", err, false))
			return
		}

		tokenType, err := jwt.GetTokenType(tokenStr)
		if err != nil || tokenType != jwt.TypeAccess {
			errHandleFunc(errsuit.NewUnauthorized("invalid token", err, false))
			return
		}

		userId, err := jwt.GetUserIdFromToken(tokenStr)
		if err != nil {
			errHandleFunc(errsuit.NewUnauthorized("invalid token payload", err, true))
			return
		}

		c.Set("user_id", userId)

		c.Next()
	}
}

func RefreshAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenPrefix := "Bearer "

		errHandleFunc := utils.MakeErrHandlerFunc(c)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, tokenPrefix) {
			errHandleFunc(errsuit.NewUnauthorized("missing authorization token", nil, false))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, tokenPrefix)
		err := jwt.ParseToken(tokenStr)
		if err != nil {
			errHandleFunc(errsuit.NewUnauthorized("invalid token", err, false))
			return
		}

		tokenType, err := jwt.GetTokenType(tokenStr)
		if err != nil || tokenType != jwt.TypeRefresh {
			errHandleFunc(errsuit.NewUnauthorized("invalid token", err, false))
			return
		}

		jti, err := jwt.GetJtiFromRefreshToken(tokenStr)
		if err != nil {
			errHandleFunc(errsuit.NewUnauthorized("invalid token payload", err, false))
			return
		}

		c.Set("jti", jti)

		userId, err := jwt.GetUserIdFromToken(tokenStr)
		if err != nil {
			errHandleFunc(errsuit.NewUnauthorized("invalid token payload", err, false))
			return
		}

		c.Set("user_id", userId)

		c.Next()
	}
}
