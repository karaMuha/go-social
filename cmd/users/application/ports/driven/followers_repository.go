package driven

import "context"

type IFollowersRepository interface {
	Follow(ctx context.Context, userID, followerID string) error
	Unfollow(ctx context.Context, followerID, userID string) error
}
