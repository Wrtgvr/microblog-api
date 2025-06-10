package postservice

import postrepo "github.com/wrtgvr/microblog/internal/repository/post"

type PostService struct {
	repo *postrepo.PostRepo
}

func NewPostService(repo postrepo.PostRepo) *PostService {
	return &PostService{
		repo: &repo,
	}
}
