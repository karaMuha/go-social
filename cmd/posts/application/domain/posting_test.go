package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreatePost(t *testing.T) {
	InitValidator()
	userID := uuid.New()
	tests := []struct {
		testName string
		title    string
		userID   string
		content  string
		wantErr  bool
	}{
		{"TestNoParams", "", "", "", true},
		{"TestNoUserIDAndContent", "SomeRandomTitle", "", "", true},
		{"TestNoTitleAndContent", "", userID.String(), "", true},
		{"TestNoTitleAndUserID", "", "", "SomeRandomContent", true},
		{"TestNoContent", "SomeRandomTitle", userID.String(), "", true},
		{"TestNoUserID", "SomeRandomTitle", "", "SomeRandomContent", true},
		{"TestNoTitle", "", userID.String(), "SomeRandomContent", true},
		{"TestSuccessfulPost", "SomeRandomTitle", userID.String(), "SomeRandomContent", false},
	}

	for _, test := range tests {
		_, err := CreatePost(test.title, test.userID, test.content)
		if err == nil && test.wantErr {
			t.Errorf("Create post test error: want error but got none for test case: %s", test.testName)
		}
		if err != nil && !test.wantErr {
			t.Errorf("Create post test error: want no error but got error for test case: %s", test.testName)
		}
	}
}
