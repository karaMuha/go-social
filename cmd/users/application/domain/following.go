package domain

import (
	"time"
)

type Following struct {
	UserID         string    `validate:"required,uuid"`
	FollowedUserID string    `validate:"required,uuid"`
	CreatedAt      time.Time `validate:"required,datetime"`
}

func Follow(userID string, followedUserID string) (*Following, error) {
	following := Following{
		UserID:         followedUserID,
		FollowedUserID: userID,
		CreatedAt:      time.Now(),
	}

	err := validate.Struct(&following)
	if err != nil {
		return nil, err
	}

	return &following, err
}

// unnecessary?
func Unfollow(following *Following, userID string) error {
	return nil
}
