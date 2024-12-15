package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/contents/application/domain"
	"github.com/karaMuha/go-social/contents/application/ports/driven"
)

type PostContentDto struct {
	Title   string `json:"title"`
	UserID  string
	Content string `json:"content"`
}

type PostContentCommand struct {
	contentsRepository driven.ContentsRepository
}

func NewPostContentCommand(contentsRepository driven.ContentsRepository) PostContentCommand {
	return PostContentCommand{
		contentsRepository: contentsRepository,
	}
}

func (c PostContentCommand) PostContent(ctx context.Context, cmd *PostContentDto) (string, error) {
	post, err := domain.PostContent(cmd.Title, cmd.UserID, cmd.Content)
	if err != nil {
		return "", fmt.Errorf("error creating post: %v", err)
	}

	postID, err := c.contentsRepository.CreateEntry(ctx, post)
	if err != nil {
		return "", fmt.Errorf("error creating post: %v", err)
	}

	return postID, nil
}
