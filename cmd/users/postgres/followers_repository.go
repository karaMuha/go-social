package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
)

type FollowersRepository struct {
	db *sql.DB
}

func NewFollowersRepository(db *sql.DB) FollowersRepository {
	return FollowersRepository{
		db: db,
	}
}

var _ driven.IFollowersRepository = (*FollowersRepository)(nil)

func (r FollowersRepository) Follow(ctx context.Context, following *domain.Following) error {
	query := `
		INSERT INTO followers (user_id, follower_id, created_at)
		VALUES($1, $2, $3)
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, following.FollowedUserID, following.UserID, following.CreatedAt)
	return err
}

func (r FollowersRepository) Unfollow(ctx context.Context, userID string, followedUserID string) error {
	query := `
		DELETE FROM followers
		WHERE user_id = $1 AND follower_id = $2
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, followedUserID, userID)
	return err
}

func (r FollowersRepository) GetFollowersOfUser(ctx context.Context, userID string) (followerList []*string, err error) {
	query := `
		SELECT follower_id
		FROM followers
		WHERE user_id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
	}(rows)

	var followerID string
	for rows.Next() {
		err = rows.Scan(&followerID)
		if err != nil {
			return
		}

		followerList = append(followerList, &followerID)
	}

	return
}
