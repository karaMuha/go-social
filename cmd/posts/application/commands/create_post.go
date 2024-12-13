package commands

import (
	"context"
	"fmt"

	"github.com/karaMuha/go-social/posts/application/domain"
	"github.com/karaMuha/go-social/posts/application/ports/driven"
)

type CreatePostDto struct {
	Title   string `json:"title"`
	UserID  string
	Content string `json:"content"`
}

type CreatePostCommand struct {
	postsRepository driven.PostsRepository
}

func NewCreatePostCommand(postsRepository driven.PostsRepository) CreatePostCommand {
	return CreatePostCommand{
		postsRepository: postsRepository,
	}
}

func (c CreatePostCommand) CreatePost(ctx context.Context, cmd *CreatePostDto) (string, error) {
	post, err := domain.CreatePost(cmd.Title, cmd.UserID, cmd.Content)
	if err != nil {
		return "", fmt.Errorf("error creating post: %v", err)
	}

	postID, err := c.postsRepository.CreateEntry(ctx, post)
	if err != nil {
		return "", fmt.Errorf("error creating post: %v", err)
	}

	return postID, nil
}
