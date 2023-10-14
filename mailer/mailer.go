package mailer

import (
	"gitlab.com/seqone/mailtick/types"
	"log"
)

// Mailer implements app.Mailer
type Mailer struct{}

type MailerInterface interface {
	Send(e types.Email) error
}

func New() *Mailer {
	return &Mailer{}
}

// Send only prints e in stdout
func (*Mailer) Send(e types.Email) error {
	log.Printf("Sending email to %s\n", e.Recipient)
	log.Printf("Subject: %s\n", e.Subject)
	log.Printf("Body: %s\n", e.Body)

	return nil
}
