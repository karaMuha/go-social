package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestFollow(t *testing.T) {
	InitValidator()
	userID := uuid.New()
	followedUserID := uuid.New()

	tests := []struct {
		testName       string
		userID         string
		followedUserID string
		wantErr        bool
	}{
		{"TestNoUserID", "", followedUserID.String(), true},
		{"TestNoFollowedUserID", userID.String(), "", true},
		{"TestUserIDSameAsFollowedUserID", userID.String(), userID.String(), true},
		{"TestSuccessfulFollow", userID.String(), followedUserID.String(), false},
	}

	for _, test := range tests {
		_, err := Follow(test.userID, test.followedUserID)
		if err == nil && test.wantErr {
			t.Errorf("Follow test error: want error but got none for test case: %s", test.testName)
		}
		if err != nil && !test.wantErr {
			t.Errorf("Follow test error: want no error but got error for test case: %s", test.testName)
		}
	}
}
