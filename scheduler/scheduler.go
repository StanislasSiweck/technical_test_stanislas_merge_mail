package scheduler

import (
	"fmt"
	"time"

	"gitlab.com/seqone/mailtick/db"
	"gitlab.com/seqone/mailtick/mailer"
)

// Scheduler holds the logic for scheduling emails send.
type Scheduler struct {
	db     db.DB
	mailer *mailer.Mailer

	stop chan struct{}
}

// New returns a new Scheduler.
func New(db db.DB, mailer *mailer.Mailer) *Scheduler {
	return &Scheduler{
		db:     db,
		mailer: mailer,
		// TODO init stop channel
	}
}

// Start starts an infinite loop. Every pause duration, it checks for pending
// emails, if found, it merges them by recipient and send them.
// In case of error during the send, the error is reported in the related
// emails.
// TODO:
//    - start a time ticker `time.NewTicker`
//    - get pending emails: `s.db.FindPendingEmails()`
//    - Merge emails for the same recipient and send one email
//    - stop the time ticker on a Stop signal. (you can use a "go channel" for that)
func (s *Scheduler) Start(pause time.Duration) error {
	return fmt.Errorf("not implemented")
}

// Stop sends the stop signal and returns once consumed or after 5 seconds.
func (s *Scheduler) Stop() {
	select {
	case <-time.After(time.Second * 5):
	case s.stop <- struct{}{}:
	}
}
