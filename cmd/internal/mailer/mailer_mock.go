package mailer

import (
	"errors"
	"strings"
)

type MailerMock struct{}

// provide "error" as value for parameter token if you want to simulate an error
// otherwise this function returns nil
func (m *MailerMock) SendRegistrationMail(email, token string) error {
	if strings.EqualFold(token, "error") {
		return errors.New("could not send email")
	}

	return nil
}
