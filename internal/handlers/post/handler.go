package posthandler

import (
	"github.com/wrtgvr/microblog/internal/handlers"
	postservice "github.com/wrtgvr/microblog/internal/services/post"
)

type PostHandler struct {
	deps    *handlers.HandlerDeps
	service *postservice.PostService
}

func NewPostHandler(deps *handlers.HandlerDeps, s *postservice.PostService) *PostHandler {
	return &PostHandler{
		deps:    deps,
		service: s,
	}
}
