package mailer

import (
	"fmt"
)

// Email has values for From/To/Subject/Body fields
type Email struct {
	FromAddress string
	ToAddress   string
	Subject     string
	Body        string
}

// Provider interface interacts with any given sender
type Provider interface {
	GetServer() string
	GetPort() int
	GetFullAddress() string
	GetUsername() string
	GetPassword() string
}

// ServerConfig has SMTP config values
type ServerConfig struct {
	Address  string
	Port     int
	Username string
	Password string
}

// GetFullAddress return full server/port as string
func (s ServerConfig) GetFullAddress() string {
	if s.Address != "" && s.Port > 0 {
		return fmt.Sprintf("%v:%v", s.Address, s.Port)
	}
	return ""
}

// GetServer returns SMTP host/server name
func (s ServerConfig) GetServer() string {
	return s.Address
}

// GetPort returns SMTP port number
func (s ServerConfig) GetPort() int {
	return s.Port
}

// GetUsername returns SMTP username credential
func (s ServerConfig) GetUsername() string {
	return s.Username
}

// GetPassword returns SMTP password credential
func (s ServerConfig) GetPassword() string {
	return s.Password
}
