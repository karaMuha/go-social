package queries

import (
	"context"

	"github.com/karaMuha/go-social/posts/application/domain"
	"github.com/karaMuha/go-social/posts/application/ports/driven"
)

type GetPostQuery struct {
	postsRepository driven.PostsRepository
}

func NewGetPostQuery(postsRepository driven.PostsRepository) GetPostQuery {
	return GetPostQuery{
		postsRepository: postsRepository,
	}
}

func (q GetPostQuery) GetPost(ctx context.Context, postID string) (*domain.Post, error) {
	return q.postsRepository.GetByID(ctx, postID)
}
