package driver

import (
	"context"

	"github.com/karaMuha/go-social/users/application/commands"
	"github.com/karaMuha/go-social/users/application/domain"
)

type IApplication interface {
	ICommands
	IQueries
}

type ICommands interface {
	SignupUser(ctx context.Context, cmd *commands.SignupUserDto) error
	ConfirmUser(ctx context.Context, cmd *commands.ConfirmUserDto) error
	FollowUser(ctx context.Context, cmd *commands.FollowUserDto) error
	UnfollowUser(ctx context.Context, cmd *commands.UnfollowUserDto) error
	ValidateUser(ctx context.Context, cmd *commands.ValidateCredentialsDto) (string, error)
}

type IQueries interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.Registration, error)
	GetUserByID(ctx context.Context, userID string) (*domain.Registration, error)
	GetFollowersOfUser(ctx context.Context, userID string) ([]*string, error)
}
