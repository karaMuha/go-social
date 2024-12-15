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
	contentsRepository driven.ContentsRepository
}

func NewUpdateContentCommand(contentsRepository driven.ContentsRepository) UpdateContentCommand {
	return UpdateContentCommand{
		contentsRepository: contentsRepository,
	}
}

func (c UpdateContentCommand) UpdateContent(ctx context.Context, cmd *UpdateContentDto) error {
	post, err := c.contentsRepository.GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	err = post.Update(cmd.Title, cmd.Content, cmd.UserID)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	err = c.contentsRepository.UpdateEntry(ctx, post)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}

	return nil
}
