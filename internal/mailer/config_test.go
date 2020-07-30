package mailer

import (
	"testing"
)

func TestServerConfigGetFullAddressEmpty(t *testing.T) {
	config := ServerConfig{}
	if config.GetFullAddress() != "" {
		t.Error("GetFullAddress was not empty")
	}
}

func TestServerConfigGetFullAddressNotEmpty(t *testing.T) {
	actual := "smtp:1234"
	config := ServerConfig{Address: "smtp", Port: 1234}
	if config.GetFullAddress() != actual {
		t.Errorf("GetFullAddress = %v; want %v", config.GetFullAddress(), actual)
	}
}

func TestServerConfigGetServerEmpty(t *testing.T) {
	config := ServerConfig{}
	if config.GetServer() != "" {
		t.Error("Server was not empty")
	}
}

func TestServerConfigGetServerNotEmpty(t *testing.T) {
	actual := "smtp"
	config := ServerConfig{Address: "smtp"}
	if config.GetServer() != actual {
		t.Errorf("GetServer = %v; want %v", config.GetServer(), actual)
	}
}

func TestServerConfigGetPortEmpty(t *testing.T) {
	config := ServerConfig{}
	if config.GetPort() != 0 {
		t.Error("Port was not empty")
	}
}
func TestServerConfigGetPortNotEmpty(t *testing.T) {
	actual := 1234
	config := ServerConfig{Port: 1234}
	if config.GetPort() != actual {
		t.Errorf("GetPort = %v; want %v", config.GetPort(), actual)
	}
}

func TestServerConfigGetUsernameEmpty(t *testing.T) {
	config := ServerConfig{}
	if config.GetUsername() != "" {
		t.Error("Username was not empty")
	}
}

func TestServerConfigGetUsernameNotEmpty(t *testing.T) {
	actual := "admin"
	config := ServerConfig{Username: "admin"}
	if config.GetUsername() != actual {
		t.Errorf("GetUsername = %v; want %v", config.GetUsername(), actual)
	}
}

func TestServerConfigGetPasswordEmpty(t *testing.T) {
	config := ServerConfig{}
	if config.GetPassword() != "" {
		t.Error("Password was not empty")
	}
}

func TestServerConfigGetPasswordNotEmpty(t *testing.T) {
	actual := "abc12345"
	config := ServerConfig{Password: "abc12345"}
	if config.GetPassword() != actual {
		t.Errorf("GetPassword = %v; want %v", config.GetPassword(), actual)
	}
}
