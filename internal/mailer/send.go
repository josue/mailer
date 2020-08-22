package mailer

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	missingServerPrefix = "missing server %v"
	missingEmailPrefix  = "missing email %v"
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
		return fmt.Errorf(missingServerPrefix, "address/port")
	}
	if e.Server.GetUsername() == "" {
		return fmt.Errorf(missingServerPrefix, "username")
	}
	if e.Server.GetPassword() == "" {
		return fmt.Errorf(missingServerPrefix, "password")
	}

	return nil
}

// validate email inputs
func (e EmailSender) validateEmail() error {
	if e.Email.FromAddress == "" {
		return fmt.Errorf(missingEmailPrefix, "FromAddress")
	}
	if e.Email.ToAddress == "" {
		return fmt.Errorf(missingEmailPrefix, "ToAddress")
	}
	if e.Email.Subject == "" {
		return fmt.Errorf(missingEmailPrefix, "Subject")
	}
	if e.Email.Body == "" {
		return fmt.Errorf(missingEmailPrefix, "Body")
	}

	return nil
}

// creates the full message payload
func (e EmailSender) buildMessage() []string {
	msg := []string{
		fmt.Sprintf("From: %v\n", e.Email.FromAddress),
		fmt.Sprintf("To: %v\n", e.Email.ToAddress),
		fmt.Sprintf("Subject: %v\n\n", e.Email.Subject),
		e.Email.Body,
	}
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
