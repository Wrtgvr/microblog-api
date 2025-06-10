package posthandler

import (
	"github.com/wrtgvr/microblog/internal/handlers"
	repo "github.com/wrtgvr/microblog/internal/repository"
	postrepo "github.com/wrtgvr/microblog/internal/repository/post"
	postservice "github.com/wrtgvr/microblog/internal/services/post"
)

func NewPostHandlerWithDeps(db *repo.Database, deps *handlers.HandlerDeps) *PostHandler {
	repo := postrepo.NewPostgresPostRepo(db)
	service := postservice.NewPostService(repo)

	return NewPostHandler(deps, service)
}
