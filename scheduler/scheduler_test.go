package scheduler

import (
	"gitlab.com/seqone/mailtick/db"
	"gitlab.com/seqone/mailtick/mailer"
	"gitlab.com/seqone/mailtick/types"
	"testing"
	"time"
)

// J'ai voulu faire des test unitaires mais je n'ai pas réussi car je n'arriverais a trouver un moyen de tester convenablement
// Je me suis tromper dans la conception de mon code, j'aurais du référencer au test en aval
func TestBasicScheduler(t *testing.T) {

	email := []types.Email{
		{
			Recipient: "bob",
			Subject:   "subject",
			Body:      "body",
		},
	}

	mockFunc := db.MockFunction{
		FindPendingEmails: func() ([]types.Email, error) {
			return email, nil
		},
		PendingEmailsToValid: func() error {
			return nil
		},
	}

	dbMock, _ := db.NewMock(mockFunc)

	scheduler := New(dbMock, mailer.New())
	go scheduler.Start(1 * time.Second)
}
