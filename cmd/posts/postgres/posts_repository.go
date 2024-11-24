package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/karaMuha/go-social/posts/application/domain"
	"github.com/karaMuha/go-social/posts/application/ports/driven"
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) PostsRepository {
	return PostsRepository{
		db: db,
	}
}

var _ driven.PostsRepository = (*PostsRepository)(nil)

func (r PostsRepository) CreateEntry(ctx context.Context, post *domain.Post) error {
	query := `
		INSERT INTO posts (title, user_id, content, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id string
	err := r.db.QueryRowContext(ctx, query, post.Title, post.UserID, post.Content, post.UpdatedAt, post.CreatedAt).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
