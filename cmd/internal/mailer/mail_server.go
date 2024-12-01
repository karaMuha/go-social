package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"time"

	"github.com/vanng822/go-premailer/premailer"
	simpleMail "github.com/xhit/go-simple-mail/v2"
)

type Mailer interface {
	SendRegistrationMail(email, token string) error
}

type Mail struct {
	From     string `json:"from"`
	FromName string `json:"fromName"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
	DataMap  map[string]any
}

type MailServer struct {
	domain      string
	host        string
	port        int
	username    string
	password    string
	encryption  string
	fromAddress string
	fromName    string
	client      *simpleMail.SMTPClient
}

var _ Mailer = (*MailServer)(nil)

func NewMailServer() *MailServer {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic(err)
	}

	mailServer := simpleMail.NewSMTPClient()
	mailServer.Host = os.Getenv("MAIL_HOST")
	mailServer.Port = port
	mailServer.Username = os.Getenv("MAIL_USERNAME")
	mailServer.Password = os.Getenv("MAIL_PASSWORD")
	mailServer.Encryption = getEncryption(os.Getenv("MAIL_ENCRYPTION"))
	mailServer.KeepAlive = true
	mailServer.ConnectTimeout = 10 * time.Second
	mailServer.SendTimeout = 10 * time.Second

	smtpClient, err := mailServer.Connect()
	if err != nil {
		panic(err)
	}

	return &MailServer{
		domain:      os.Getenv("MAIL_DOMAIN"),
		host:        os.Getenv("MAIL_HOST"),
		port:        port,
		username:    os.Getenv("MAIL_USERNAME"),
		password:    os.Getenv("MAIL_PASSWORD"),
		encryption:  os.Getenv("MAIL_ENCRYPTION"),
		fromAddress: os.Getenv("MAIL_FROMADDRESS"),
		fromName:    os.Getenv("MAIL_FROMNAME"),
		client:      smtpClient,
	}
}

func (m *MailServer) SendRegistrationMail(address, token string) error {
	mailMessage := fmt.Sprintf("Please visit localhost:8080/v1/users/confirm?email=%s&token=%s to complete your registration", address, token)

	mail := Mail{
		To:      address,
		Subject: "Confirm Registration",
		Message: mailMessage,
	}

	if mail.From == "" {
		mail.From = m.fromAddress
	}

	if mail.FromName == "" {
		mail.FromName = m.fromName
	}

	data := map[string]any{
		"message": mail.Message,
	}

	mail.DataMap = data

	formattedMessage, err := m.buildHTMLMail(&mail)

	if err != nil {
		return err
	}

	email := simpleMail.NewMSG()
	email.SetFrom(mail.From)
	email.AddTo(mail.To)
	email.SetSubject(mail.Subject)
	email.SetBody(simpleMail.TextHTML, formattedMessage)

	err = email.Error

	if err != nil {
		return fmt.Errorf("error creating confirmation mail: %v", err)
	}

	err = email.Send(m.client)

	if err != nil {
		return fmt.Errorf("error sending confirmation mail: %v", err)
	}

	return nil
}

func (m *MailServer) buildHTMLMail(mail *Mail) (string, error) {
	templateToRender := "./app/mail.html"

	t, err := template.New("email-html").ParseFiles(templateToRender)

	if err != nil {
		return "", err
	}

	var template bytes.Buffer
	err = t.ExecuteTemplate(&template, "body", mail.DataMap)

	if err != nil {
		return "", err
	}

	formattedMessage := template.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)

	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *MailServer) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)

	if err != nil {
		return "", err
	}

	html, err := prem.Transform()

	if err != nil {
		return "", err
	}

	return html, nil
}

func getEncryption(s string) simpleMail.Encryption {
	switch s {
	case "tls":
		return simpleMail.EncryptionSTARTTLS
	case "ssl":
		return simpleMail.EncryptionSSLTLS
	case "none", "":
		return simpleMail.EncryptionNone
	default:
		return simpleMail.EncryptionSTARTTLS
	}
}
