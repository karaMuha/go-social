package driven

import (
	"context"

	"github.com/karaMuha/go-social/users/application/domain"
)

type IFollowersRepository interface {
	Follow(ctx context.Context, following *domain.Following) error
	Unfollow(ctx context.Context, userID string, followedUserID string) error
	GetFollowersOfUser(ctx context.Context, followedID string) ([]*string, error)
}
