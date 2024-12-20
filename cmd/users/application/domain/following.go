package domain

import (
	"errors"
	"time"
)

type Following struct {
	UserID         string `validate:"required,uuid"`
	FollowedUserID string `validate:"required,uuid"`
	CreatedAt      time.Time
}

func Follow(userID string, followedUserID string) (*Following, error) {
	if userID == followedUserID {
		return nil, errors.New("you cannot follow yourself")
	}

	following := Following{
		UserID:         userID,
		FollowedUserID: followedUserID,
	}

	err := validate.Struct(&following)
	if err != nil {
		return nil, err
	}

	following.CreatedAt = time.Now()

	return &following, err
}

// unnecessary?
func Unfollow(following *Following, userID string) error {
	return nil
}
