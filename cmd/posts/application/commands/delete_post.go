package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/posts/application/ports/driven"
)

type DeletePostDto struct {
	ID     string
	UserID string
}

type DeletePostCommand struct {
	postsRepository driven.PostsRepository
}

func NewDeletePostCommand(postsRepository driven.PostsRepository) DeletePostCommand {
	return DeletePostCommand{
		postsRepository: postsRepository,
	}
}

func (c DeletePostCommand) DeletePost(ctx context.Context, cmd *DeletePostDto) error {
	post, err := c.postsRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("error deleting post: %v", err)
	}

	err = post.Delete(cmd.UserID)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	err = c.postsRepository.DeleteEntry(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	return nil
}
