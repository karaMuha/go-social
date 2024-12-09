package domain

import (
	"testing"

	"github.com/thanhpk/randstr"
)

func TestSignup(t *testing.T) {
	InitValidator()
	tests := []struct {
		testName string
		userName string
		email    string
		password string
		wantErr  bool
	}{
		{"TestNoParameters", "", "", "", true},
		{"TestNoEmailAndPassword", "John", "", "", true},
		{"TestNoUsernameAndPassword", "", "test@test.com", "", true},
		{"TestNoUsernameAndEmail", "", "", "test123", true},
		{"TestNoPassword", "John", "test@test.com", "", true},
		{"TestNoEmail", "John", "", "test123", true},
		{"TestNoUsername", "", "test@test.com", "test123", true},
		{"TestSuccessfulSignup", "John", "test@test.com", "test123", false},
	}

	for _, test := range tests {
		registration, err := Signup(test.userName, test.email, test.password)
		if err == nil && test.wantErr {
			t.Errorf("Signup test error: want error but got none for test case: %s", test.testName)
		}
		if err != nil && !test.wantErr {
			t.Errorf("Signup test error: want no error but got error for test case: %s", test.testName)
		}
		if err == nil {
			if registration.Username != test.userName {
				t.Errorf("Signup test error: got %s but want %s for test case %s", registration.Username, test.userName, test.testName)
			}
			if registration.Email != test.email {
				t.Errorf("Signup test error: got %s but want %s for test case %s", registration.Email, test.email, test.testName)
			}
			if registration.Active {
				t.Errorf("Signup test error: registration active is true but want false for test case %v", test.testName)
			}
			if registration.RegistrationToken == "" {
				t.Errorf("Signup test error: signup successful but got no registration token for test case %s", test.testName)
			}
		}
	}
}

func TestActivate(t *testing.T) {
	tokenFromDb := randstr.String(REGISTRATION_TOKEN_LENGTH)
	tokenFromRequest := tokenFromDb
	tests := []struct {
		testName         string
		active           bool
		tokenFromDb      string
		tokenFromRequest string
		wantErr          bool
	}{
		{"TestUserAlreadyActive", true, tokenFromDb, tokenFromRequest, true},
		{"TestTokenFromDbWrongLength", false, "tooShortToken", tokenFromRequest, true},
		{"TestTokenDoNotMatch", false, tokenFromDb, "wrongToken", true},
		{"TestSuccessfulActivation", false, tokenFromDb, tokenFromRequest, false},
	}

	for _, test := range tests {
		err := Activate(test.active, test.tokenFromDb, test.tokenFromRequest)
		if err == nil && test.wantErr {
			t.Errorf("Signup test error: want error but got none for test case: %s", test.testName)
		}
		if err != nil && !test.wantErr {
			t.Errorf("Signup test error: want no error but got error for test case: %s", test.testName)
		}
	}
}
