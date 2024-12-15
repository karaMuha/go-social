package queries

import (
	"context"

	"github.com/karaMuha/go-social/contents/application/domain"
	"github.com/karaMuha/go-social/contents/application/ports/driven"
)

type GetContentDetailsQuery struct {
	postsRepository driven.PostsRepository
}

func NewGetContentDetailsQuery(postsRepository driven.PostsRepository) GetContentDetailsQuery {
	return GetContentDetailsQuery{
		postsRepository: postsRepository,
	}
}

func (q GetContentDetailsQuery) GetContentDetails(ctx context.Context, postID string) (*domain.Content, error) {
	return q.postsRepository.GetByID(ctx, postID)
}
