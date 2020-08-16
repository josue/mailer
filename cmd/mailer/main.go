package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/josue/mailer/internal/mailer"
)

func main() {
	// ENV params
	serverUsername := os.Getenv("MAILER_USERNAME")
	serverPassword := os.Getenv("MAILER_PASSWORD")

	// CLI params
	fromAddr := flag.String("from", "", "from address")
	toAddr := flag.String("to", "", "to address")
	subject := flag.String("subject", "Hello Friend", "subject")
	body := flag.String("body", "Are you still there?", "body")
	flag.Parse()

	timeStart := time.Now()

	// construct email params
	email := mailer.Email{
		FromAddress: *fromAddr,
		ToAddress:   *toAddr,
		Subject:     *subject,
		Body:        *body,
	}

	// construct server params
	var username string
	if serverUsername != "" {
		username = serverUsername
	} else {
		username = email.FromAddress
	}
	server := mailer.NewGmail(
		mailer.ServerConfig{
			Username: username,
			Password: serverPassword,
		},
	)

	// construct sender params
	sender := mailer.NewEmailSender()
	sender.Server = server
	sender.Email = email

	// send email to given sender server
	err := sender.SendMessage()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Printf("[Elapsed: %v] Success, sent from (%v) to (%v) with SMTP server (%v)\n", time.Since(timeStart), email.FromAddress, email.ToAddress, server.GetServer())
}
