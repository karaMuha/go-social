package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/karaMuha/go-social/contents/application/domain"
	"github.com/karaMuha/go-social/contents/application/ports/driven"
)

type ContentsRepository struct {
	db *sql.DB
}

func NewContentsRepository(db *sql.DB) ContentsRepository {
	return ContentsRepository{
		db: db,
	}
}

var _ driven.PostsRepository = (*ContentsRepository)(nil)

func (r ContentsRepository) CreateEntry(ctx context.Context, content *domain.Content) (string, error) {
	query := `
		INSERT INTO posts (title, user_id, content, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id string
	err := r.db.QueryRowContext(ctx, query, content.Title, content.UserID, content.Infill, content.UpdatedAt, content.CreatedAt).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r ContentsRepository) GetByID(ctx context.Context, contentID string) (*domain.Content, error) {
	query := `
		SELECT *
		FROM posts
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var content domain.Content
	err := r.db.QueryRowContext(ctx, query, contentID).Scan(
		&content.ID,
		&content.Title,
		&content.UserID,
		&content.Infill,
		&content.UpdatedAt,
		&content.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (r ContentsRepository) UpdateEntry(ctx context.Context, content *domain.Content) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, updated_at = $3
		WHERE id = $4
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, content.Title, content.Infill, content.UpdatedAt, content.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r ContentsRepository) DeleteEntry(ctx context.Context, contentID string) error {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, contentID)
	if err != nil {
		return err
	}

	return nil
}
