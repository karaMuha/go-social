package queries

import (
	"context"

	"github.com/karaMuha/go-social/contents/application/domain"
	"github.com/karaMuha/go-social/contents/application/ports/driven"
)

type GetContentOfUserQuery struct {
	contentsRepository driven.ContentsRepository
}

func NewGetContentOfUserQuery(contentsRepository driven.ContentsRepository) GetContentOfUserQuery {
	return GetContentOfUserQuery{
		contentsRepository: contentsRepository,
	}
}

func (q GetContentOfUserQuery) GetContentOfUser(ctx context.Context, userID string) ([]*domain.Content, error) {
	return q.contentsRepository.GetAllOfUser(ctx, userID)
}
