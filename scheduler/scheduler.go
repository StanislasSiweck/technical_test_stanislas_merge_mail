package scheduler

import (
	"gitlab.com/seqone/mailtick/types"
	"log"
	"time"

	"gitlab.com/seqone/mailtick/db"
	"gitlab.com/seqone/mailtick/mailer"
)

// Scheduler holds the logic for scheduling emails send.
type Scheduler struct {
	db     db.DB
	mailer *mailer.Mailer

	stop time.Ticker
}

// New returns a new Scheduler.
func New(db db.DB, mailer *mailer.Mailer) *Scheduler {
	return &Scheduler{
		db:     db,
		mailer: mailer,
	}
}

// Start starts an infinite loop. Every pause duration, it checks for pending
// emails, if found, it merges them by recipient and send them.
// In case of error during the send, the error is reported in the related
// emails.
// TODO:
//   - start a time ticker `time.NewTicker`
//   - get pending emails: `s.db.FindPendingEmails()`
//   - Merge emails for the same recipient and send one email
//   - stop the time ticker on a Stop signal. (you can use a "go channel" for that)

// Autre solution aurais été d'utiliser un cron avec github.com/robfig/cron
func (s *Scheduler) Start(pause time.Duration) {
	s.stop = *time.NewTicker(pause)

	for range s.stop.C {
		emails, err := s.db.FindPendingEmails()
		log.Printf("Start merge and send %v emails", len(emails))
		if err != nil {
			continue
		}

		emailsByRecipient := make(map[string]types.Email)
		for _, email := range emails {

			emailByRecipient, ok := emailsByRecipient[email.Recipient]
			if ok {
				// utilisation des tr dans l'idée de faire une liste
				emailByRecipient.Body += "<tr>" + email.Body + "</tr>"
				emailByRecipient.Subject += ";" + email.Subject
			} else {
				email.Body = "<tr>" + email.Body + "</tr>"
				emailByRecipient = email
			}

			emailsByRecipient[email.Recipient] = emailByRecipient
		}

		for _, email := range emailsByRecipient {
			err = s.mailer.Send(email)
			if err != nil {
				if err = s.db.PendingEmailsToError(err, email.Recipient); err != nil {
					log.Printf("Error while updating email status to error: %v", err)
				}
			}
		}

		if err = s.db.PendingEmailsToValid(); err != nil {
			log.Printf("Error while updating email status to valid: %v", err)
		}
		s.stop.Reset(pause)
	}
}
