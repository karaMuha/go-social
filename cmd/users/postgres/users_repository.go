package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return UsersRepository{
		db: db,
	}
}

var _ driven.IUsersRepsitory = (*UsersRepository)(nil)

func (r UsersRepository) CreateEntry(ctx context.Context, registration *domain.Registration) error {
	query := `
		INSERT INTO users (email, username, user_password, registration_token, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id string
	err := r.db.QueryRowContext(
		ctx,
		query,
		registration.Email,
		registration.Username,
		registration.Password,
		registration.RegistrationToken,
		registration.CreatedAt,
	).Scan(&id)

	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			if strings.Contains(err.Error(), "email") {
				return errors.New("email already exists")
			}
			if strings.Contains(err.Error(), "username") {
				return errors.New("username already exists")
			}
		}
		return err
	}

	return nil
}

func (r UsersRepository) GetByID(ctx context.Context, userID string) (*domain.Registration, error) {
	query := `
		SELECT * FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.Registration
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r UsersRepository) GetByEmail(ctx context.Context, email string) (*domain.Registration, error) {
	query := `
		SELECT * FROM users
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.Registration
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r UsersRepository) DeleteEntry(ctx context.Context, userID string) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if _, err := r.db.ExecContext(ctx, query, userID); err != nil {
		return err
	}

	return nil
}
