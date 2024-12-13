package mailer

type IMailer interface {
	SendRegistrationMail(email, token string) error
}
