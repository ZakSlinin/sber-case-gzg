package mailer

import (
	"fmt"
	"net/smtp"
)

type EmailVerificationMailer interface {
	Send(to, subject, body string) error
}

type SMTPMailer struct {
	host     string
	port     string
	username string
	password string
}

func NewEmailVerificationMailer(host, port, username, password string) *SMTPMailer {
	return &SMTPMailer{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (m *SMTPMailer) Send(to, subject, body string) error {
	address := m.host + ":" + m.port
	auth := smtp.PlainAuth("Sber", m.username, m.password, m.host)

	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body)
	err := smtp.SendMail(address, auth, m.username, []string{to}, []byte(message))

	return err
}
