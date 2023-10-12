package mailer

import (
	"gitlab.com/seqone/mailtick/types"
)

// Mailer implements app.Mailer
type Mailer struct{}

func New() *Mailer {
	return &Mailer{}
}

// Send only prints e in stdout
func (*Mailer) Send(e types.Email) error {
	// TODO it's not mandatory to implement a mail server.
	// You can just simulate it through a log.
	return nil
}
