package mailer

import (
	"errors"
	"strings"
)

type MailerMock struct{}

// provide "error@error.com" as value for parameter "email" if you want to simulate an error
// otherwise this function returns nil
func (m *MailerMock) SendRegistrationMail(email, token string) error {
	if strings.EqualFold(email, "error@error.com") {
		return errors.New("could not send email")
	}

	return nil
}
