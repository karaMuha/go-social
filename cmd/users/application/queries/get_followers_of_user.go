package queries

import (
	"context"

	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type GetFollowerQuery struct {
	followersRepository driven.IFollowersRepository
}

func NewGetFollowerQuery(followersRepository driven.IFollowersRepository) GetFollowerQuery {
	return GetFollowerQuery{
		followersRepository: followersRepository,
	}
}

func (q GetFollowerQuery) GetFollowersOfUser(ctx context.Context, userID string) ([]*string, error) {
	return q.followersRepository.GetFollowersOfUser(ctx, userID)
}
