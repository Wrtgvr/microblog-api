package authhandler

import (
	"github.com/wrtgvr/microblog/internal/handlers"
	authservice "github.com/wrtgvr/microblog/internal/services/auth"
)

type AuthHandler struct {
	deps    *handlers.HandlerDeps
	service *authservice.AuthService
}

func NewAuthHandler(deps *handlers.HandlerDeps, s *authservice.AuthService) *AuthHandler {
	return &AuthHandler{
		deps:    deps,
		service: s,
	}
}
