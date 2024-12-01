package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type FollowUserDto struct {
	UserID         string `json:"user_id"`
	FollowedUserID string `json:"followed_user_id"`
}

type FollowUserCommand struct {
	followersRepository driven.IFollowersRepository
}

func NewFollowUserCommand(followersRepository driven.IFollowersRepository) FollowUserCommand {
	return FollowUserCommand{
		followersRepository: followersRepository,
	}
}

func (c FollowUserCommand) FollowUser(ctx context.Context, cmd *FollowUserDto) error {
	following, err := domain.Follow(cmd.UserID, cmd.FollowedUserID)
	if err != nil {
		return fmt.Errorf("error following user: %v", err)
	}

	err = c.followersRepository.Follow(ctx, following)
	if err != nil {
		return fmt.Errorf("error following user: %v", err)
	}

	return nil
}
