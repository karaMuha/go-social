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

func TestUpdatePost(t *testing.T) {
	InitValidator()
	userID := uuid.New()
	wrongUserID := uuid.New()
	post, err := CreatePost("This is a Title", userID.String(), "This is the content")
	if err != nil {
		t.Errorf("Cannot prepare post for test: %v", err)
	}

	tests := []struct {
		testName string
		userID   string
		title    string
		content  string
		wantErr  bool
	}{
		{"TestNoParameters", "", "", "", true},
		{"TestWrongUser", wrongUserID.String(), "SomeRandomTitle", "SomeRandomContent", true},
		{"TestNoTitle", userID.String(), "", "SomeRandomContent", true},
		{"TestNoContent", userID.String(), "SomeRandomTitle", "", true},
		{"TestSuccessfulUpdate", userID.String(), "SomeRandomTitle", "SomeRandomContent", false},
	}

	for _, test := range tests {
		err = post.Update(test.title, test.content, test.userID)
		if err == nil && test.wantErr {
			t.Errorf("Update post test error: want error but got none for test case: %s", test.testName)
		}
		if err != nil && !test.wantErr {
			t.Errorf("Update post test error: want no error but got error for test case: %s", test.testName)
		}
	}
}

func TestDeletePost(t *testing.T) {
	InitValidator()
	userID := uuid.New()
	wrongUserID := uuid.New()
	post, err := CreatePost("This is a Title", userID.String(), "This is the content")
	if err != nil {
		t.Errorf("Cannot prepare post for test: %v", err)
	}

	tests := []struct {
		testName string
		userID   string
		wantErr  bool
	}{
		{"TestWrongUserID", wrongUserID.String(), true},
		{"TestSuccessfulDelete", userID.String(), false},
	}

	for _, test := range tests {
		err = post.Delete(test.userID)
		if err == nil && test.wantErr {
			t.Errorf("Delete post test error: want error but got none for test case: %s", test.testName)
		}
		if err != nil && !test.wantErr {
			t.Errorf("Delete post test error: want no error but got error for test case: %s", test.testName)
		}
	}
}
