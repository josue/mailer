package mailer

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

// EmailSender struct to construct a mailer instance with server/email inputs
type EmailSender struct {
	SendMail func(string, smtp.Auth, string, []string, []byte) error
	Server   Provider
	Email    Email
}

// NewEmailSender returns EmailSender struct with smtp.Sendmail
func NewEmailSender() *EmailSender {
	return &EmailSender{SendMail: smtp.SendMail}
}

// validate server inputs
func (e EmailSender) validateServer() error {
	if e.Server.GetFullAddress() == "" {
		return errors.New("Missing server address/port")
	}
	if e.Server.GetUsername() == "" {
		return errors.New("Missing server username")
	}
	if e.Server.GetPassword() == "" {
		return errors.New("Missing server password")
	}

	return nil
}

// validate email inputs
func (e EmailSender) validateEmail() error {
	if e.Email.FromAddress == "" {
		return errors.New("Missing email FromAddress")
	}
	if e.Email.ToAddress == "" {
		return errors.New("Missing email ToAddress")
	}
	if e.Email.Subject == "" {
		return errors.New("Missing email Subject")
	}
	if e.Email.Body == "" {
		return errors.New("Missing email Body")
	}

	return nil
}

// creates the full message payload
func (e EmailSender) buildMessage() []string {
	msg := []string{}
	msg = append(msg, fmt.Sprintf("From: %v\n", e.Email.FromAddress))
	msg = append(msg, fmt.Sprintf("To: %v\n", e.Email.ToAddress))
	msg = append(msg, fmt.Sprintf("Subject: %v\n\n", e.Email.Subject))
	msg = append(msg, e.Email.Body)
	return msg
}

// SendMessage constructs inputs from server/email and send the email message or returns error
func (e EmailSender) SendMessage() error {
	// validate server
	err := e.validateServer()
	if err != nil {
		return err
	}
	// validate email
	err = e.validateEmail()
	if err != nil {
		return err
	}

	// construct message
	msg := e.buildMessage()

	// send via SMTP provider
	err = e.SendMail(
		e.Server.GetFullAddress(),
		smtp.PlainAuth("", e.Server.GetUsername(), e.Server.GetPassword(), e.Server.GetServer()),
		e.Email.FromAddress,
		[]string{e.Email.ToAddress},
		[]byte(strings.Join(msg, "")),
	)

	if err != nil {
		return err
	}

	return nil
}
