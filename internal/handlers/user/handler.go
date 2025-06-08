package userhandler

import (
	"github.com/wrtgvr/microblog/internal/handlers"
	userservice "github.com/wrtgvr/microblog/internal/services/user"
)

type UserHandler struct {
	deps    *handlers.HandlerDeps
	service *userservice.UserService
}

func NewUserHandler(deps *handlers.HandlerDeps, s *userservice.UserService) *UserHandler {
	return &UserHandler{
		deps:    deps,
		service: s,
	}
}
