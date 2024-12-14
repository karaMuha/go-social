package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/karaMuha/go-social/contents/application/domain"
	"github.com/karaMuha/go-social/contents/application/ports/driven"
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

func (r PostsRepository) CreateEntry(ctx context.Context, post *domain.Post) (string, error) {
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
		return "", err
	}

	return id, nil
}

func (r PostsRepository) GetByID(ctx context.Context, postID string) (*domain.Post, error) {
	query := `
		SELECT *
		FROM posts
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var post domain.Post
	err := r.db.QueryRowContext(ctx, query, postID).Scan(
		&post.ID,
		&post.Title,
		&post.UserID,
		&post.Content,
		&post.UpdatedAt,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r PostsRepository) UpdateEntry(ctx context.Context, post *domain.Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, updated_at = $3
		WHERE id = $4
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, post.Title, post.Content, post.UpdatedAt, post.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r PostsRepository) DeleteEntry(ctx context.Context, postID string) error {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}

	return nil
}
