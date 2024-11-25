package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/karaMuha/go-social/posts/application/ports/driven"
)

type UpdatePostDto struct {
	ID      string
	Title   string
	UserID  string
	Content string
}

type UpdatePostCommand struct {
	postsRepository driven.PostsRepository
}

func NewUpdatePostCommand(postsRepository driven.PostsRepository) UpdatePostCommand {
	return UpdatePostCommand{
		postsRepository: postsRepository,
	}
}

func (c UpdatePostCommand) UpdatePost(ctx context.Context, cmd *UpdatePostDto) error {
	post, err := c.postsRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	if cmd.UserID != post.UserID {
		return errors.New("error updating post: insufficient permission")
	}

	post.Update(cmd.Title, cmd.Content)
	return nil
}
