package mailer

type Mailer interface {
	SendRegistrationMail(email, token string) error
}
