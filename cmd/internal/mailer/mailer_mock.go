package mailer

import "errors"

type SuccessMailer struct{}

func (m *SuccessMailer) SendRegistrationMail(email, token string) error {
	return nil
}

type FailMailer struct{}

func (m *FailMailer) SendRegistrationMail(email, token string) error {
	return errors.New("something went wrong")
}
