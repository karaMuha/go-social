package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/contents/application/ports/driven"
)

type RemoveContentDto struct {
	ID     string `json:"id"`
	UserID string
}

type RemoveContentCommand struct {
	postsRepository driven.PostsRepository
}

func NewRemoveContentCommand(postsRepository driven.PostsRepository) RemoveContentCommand {
	return RemoveContentCommand{
		postsRepository: postsRepository,
	}
}

func (c RemoveContentCommand) RemoveContent(ctx context.Context, cmd *RemoveContentDto) error {
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
