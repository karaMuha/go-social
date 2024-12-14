package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/contents/application/ports/driven"
)

type UpdateContentDto struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	UserID  string
	Content string `json:"content"`
}

type UpdateContentCommand struct {
	postsRepository driven.PostsRepository
}

func NewUpdateContentCommand(postsRepository driven.PostsRepository) UpdateContentCommand {
	return UpdateContentCommand{
		postsRepository: postsRepository,
	}
}

func (c UpdateContentCommand) UpdateContent(ctx context.Context, cmd *UpdateContentDto) error {
	post, err := c.postsRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	err = post.Update(cmd.Title, cmd.Content, cmd.UserID)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	err = c.postsRepository.UpdateEntry(ctx, post)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	return nil
}
