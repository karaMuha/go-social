package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type UnfollowUserDto struct {
	UserID         string `json:"user_id"`
	FollowedUserID string `json:"followed_user_id"`
}

type UnfollowUserCommand struct {
	followersRepository driven.IFollowersRepository
}

func NewUnfollowUserCommand(followersRepository driven.IFollowersRepository) UnfollowUserCommand {
	return UnfollowUserCommand{
		followersRepository: followersRepository,
	}
}

func (c UnfollowUserCommand) UnfollowUser(ctx context.Context, cmd *UnfollowUserDto) error {
	err := c.followersRepository.Unfollow(ctx, cmd.UserID, cmd.FollowedUserID)
	if err != nil {
		return fmt.Errorf("error unfollowing user: %v", err)
	}

	return nil
}
