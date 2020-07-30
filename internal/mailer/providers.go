package mailer

// Gmail struct specific to Gmail SMTP
type Gmail struct {
	ServerConfig
}

// Yahoo struct specific to Yahoo SMTP
type Yahoo struct {
	ServerConfig
}

// NewGmail returns Gmail struct with destination address/port
func NewGmail(s ServerConfig) *Gmail {
	p := &Gmail{s}
	p.Address = "smtp.gmail.com"
	p.Port = 587
	return p
}

// NewYahoo returns Yahoo struct with destination address/port
func NewYahoo(s ServerConfig) *Yahoo {
	p := &Yahoo{s}
	p.Address = "smtp.mail.yahoo.com"
	p.Port = 465 // 465 or 587
	return p
}
