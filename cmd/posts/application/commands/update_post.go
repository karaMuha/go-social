package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/posts/application/ports/driven"
)

type UpdatePostDto struct {
	ID      string
	Title   string `json:"title"`
	UserID  string
	Content string `json:"content"`
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
