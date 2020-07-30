package mailer

import (
	"net/smtp"
	"strings"
	"testing"
)

var assumeServerConfig = ServerConfig{
	Address:  "smtp",
	Port:     1234,
	Username: "a",
	Password: "b",
}

func TestSendMessageMissingServerAddressPort(t *testing.T) {
	email := Email{}
	server := ServerConfig{}
	sender := EmailSender{Server: server, Email: email}
	err := sender.SendMessage()

	if !strings.Contains(err.Error(), "Missing server address/port") {
		t.Errorf("Should be missing server address/port, got instead %v", err)
	}
}

func TestSendMessageMissingServerUsername(t *testing.T) {
	email := Email{}
	server := ServerConfig{Address: "smtp", Port: 1234}
	sender := EmailSender{Server: server, Email: email}
	err := sender.SendMessage()

	if !strings.Contains(err.Error(), "Missing server username") {
		t.Errorf("Should be missing server username, got instead %v", err)
	}
}

func TestSendMessageMissingServerPassword(t *testing.T) {
	email := Email{}
	server := ServerConfig{Address: "smtp", Port: 1234, Username: "a"}
	sender := EmailSender{Server: server, Email: email}
	err := sender.SendMessage()

	if !strings.Contains(err.Error(), "Missing server password") {
		t.Errorf("Should be missing server password, got instead %v", err)
	}
}

func TestSendMessageMissingFromAddress(t *testing.T) {
	email := Email{}
	server := assumeServerConfig
	sender := EmailSender{Server: server, Email: email}
	err := sender.SendMessage()

	if !strings.Contains(err.Error(), "Missing email FromAddress") {
		t.Errorf("Should be missing email.FromAddress, got instead %v", err)
	}
}

func TestSendMessageMissingToAddress(t *testing.T) {
	email := Email{FromAddress: "a"}
	server := assumeServerConfig
	sender := EmailSender{Server: server, Email: email}
	err := sender.SendMessage()

	if !strings.Contains(err.Error(), "Missing email ToAddress") {
		t.Errorf("Should be missing email ToAddress, got instead %v", err)
	}
}

func TestSendMessageMissingSubject(t *testing.T) {
	email := Email{
		FromAddress: "a",
		ToAddress:   "b",
	}
	server := assumeServerConfig
	sender := EmailSender{Server: server, Email: email}
	err := sender.SendMessage()

	if !strings.Contains(err.Error(), "Missing email Subject") {
		t.Errorf("Should be missing email Subject, got instead %v", err)
	}
}

func TestSendMessageMissingBody(t *testing.T) {
	email := Email{
		FromAddress: "a",
		ToAddress:   "b",
		Subject:     "hello",
	}
	server := assumeServerConfig
	sender := EmailSender{Server: server, Email: email}
	err := sender.SendMessage()

	if !strings.Contains(err.Error(), "Missing email Body") {
		t.Errorf("Should be missing email Body, got instead %v", err)
	}
}

func TestSendMessageFailNoHostExist(t *testing.T) {
	email := Email{
		FromAddress: "alice@example.com",
		ToAddress:   "bob@example.com",
		Subject:     "fsociety",
		Body:        "LEAVE ME HERE",
	}
	server := assumeServerConfig
	sender := NewEmailSender()
	sender.Server = server
	sender.Email = email
	err := sender.SendMessage()

	if !strings.Contains(err.Error(), "no such host") {
		t.Errorf("SMTP error expected; received %v", err)
	}
}

func TestSendMessageSendSuccess(t *testing.T) {
	email := Email{
		FromAddress: "alice@example.com",
		ToAddress:   "bob@example.com",
		Subject:     "fsociety",
		Body:        "LEAVE ME HERE",
	}
	server := NewGmail(
		ServerConfig{
			Username: email.FromAddress,
			Password: "secret",
		},
	)

	sender := EmailSender{}
	sender.Server = server
	sender.Email = email
	sender.SendMail = func(string, smtp.Auth, string, []string, []byte) error {
		return nil
	}

	err := sender.SendMessage()

	if err != nil {
		t.Errorf("SMTP success expected; received error instead: %v", err)
	}
}
