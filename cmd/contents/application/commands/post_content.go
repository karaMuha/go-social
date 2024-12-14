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
	postsRepository driven.PostsRepository
}

func NewPostContentCommand(postsRepository driven.PostsRepository) PostContentCommand {
	return PostContentCommand{
		postsRepository: postsRepository,
	}
}

func (c PostContentCommand) PostContent(ctx context.Context, cmd *PostContentDto) (string, error) {
	post, err := domain.PostContent(cmd.Title, cmd.UserID, cmd.Content)
	if err != nil {
		return "", fmt.Errorf("error creating post: %v", err)
	}

	postID, err := c.postsRepository.CreateEntry(ctx, post)
	if err != nil {
		return "", fmt.Errorf("error creating post: %v", err)
	}

	return postID, nil
}
