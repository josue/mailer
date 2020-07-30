package mailer

import (
	"testing"
)

func TestNewGmailCorrectServer(t *testing.T) {
	address := "smtp.gmail.com"
	server := NewGmail(ServerConfig{})
	if server.GetServer() != address {
		t.Errorf("Server = %v; want %v", server.GetServer(), address)
	}
}

func TestNewGmailWrongServer(t *testing.T) {
	address := "example"
	server := NewGmail(ServerConfig{})
	if server.GetServer() == address {
		t.Errorf("Server = %v; want %v", server.GetServer(), address)
	}
}

func TestNewYahooCorrectServer(t *testing.T) {
	address := "smtp.mail.yahoo.com"
	server := NewYahoo(ServerConfig{})
	if server.GetServer() != address {
		t.Errorf("Server = %v; want %v", server.GetServer(), address)
	}
}

func TestNewYahooWrongServer(t *testing.T) {
	address := "example"
	server := NewYahoo(ServerConfig{})
	if server.GetServer() == address {
		t.Errorf("Server = %v; want %v", server.GetServer(), address)
	}
}
