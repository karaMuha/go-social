package queries

import (
	"context"

	"github.com/karaMuha/go-social/contents/application/domain"
	"github.com/karaMuha/go-social/contents/application/ports/driven"
)

type GetContentDetailsQuery struct {
	contentsRepository driven.ContentsRepository
}

func NewGetContentDetailsQuery(postsRepository driven.ContentsRepository) GetContentDetailsQuery {
	return GetContentDetailsQuery{
		contentsRepository: postsRepository,
	}
}

func (q GetContentDetailsQuery) GetContentDetails(ctx context.Context, postID string) (*domain.Content, error) {
	return q.contentsRepository.GetByID(ctx, postID)
}
