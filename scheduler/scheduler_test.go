package scheduler

import (
	"errors"
	"gitlab.com/seqone/mailtick/db"
	"gitlab.com/seqone/mailtick/mailer"
	"gitlab.com/seqone/mailtick/types"
	"testing"
)

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
	err := scheduler.MergeAndSendEmails()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestNoMailsInDb(t *testing.T) {

	//No emails in db
	mockFunc := db.MockFunction{
		FindPendingEmails: func() ([]types.Email, error) {
			return nil, errors.New(types.ErrNoPendingEmails)
		},
	}
	dbMock, _ := db.NewMock(mockFunc)

	scheduler := New(dbMock, mailer.New())
	if err := scheduler.MergeAndSendEmails(); err == nil {
		t.Errorf("Error: %s", err)
	}

}

func TestErrorToSendEmail(t *testing.T) {
	emails := []types.Email{
		{
			Recipient: "bob",
			Subject:   "subject",
			Body:      "body",
		},
	}

	mailerMock := mailer.MockFunction{
		Send: func(e types.Email) error {
			return errors.New("error")
		},
	}
	mailer := mailer.NewMock(mailerMock)

	dbMock := db.MockFunction{
		FindPendingEmails: func() ([]types.Email, error) {
			return emails, nil
		},
		PendingEmailsToError: func(err error, recipient string) error {
			return nil
		},
		PendingEmailsToValid: func() error {
			return nil
		},
	}
	db, _ := db.NewMock(dbMock)

	scheduler := New(db, mailer)
	if err := scheduler.MergeAndSendEmails(); err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestErrorToSendEmailAndDb(t *testing.T) {
	emails := []types.Email{
		{
			Recipient: "bob",
			Subject:   "subject",
			Body:      "body",
		},
	}

	mailerMock := mailer.MockFunction{
		Send: func(e types.Email) error {
			return errors.New("error")
		},
	}
	mailer := mailer.NewMock(mailerMock)

	dbMock := db.MockFunction{
		FindPendingEmails: func() ([]types.Email, error) {
			return emails, nil
		},
		PendingEmailsToError: func(err error, recipient string) error {
			return errors.New("can not update status to error")
		},
	}
	db, _ := db.NewMock(dbMock)

	scheduler := New(db, mailer)
	if err := scheduler.MergeAndSendEmails(); err == nil && err.Error() != "can not update status to error" {
		t.Errorf("Error: %s", err)
	}
}

func TestErrorToValid(t *testing.T) {

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
			return errors.New("can not update status to valid")
		},
	}

	dbMock, _ := db.NewMock(mockFunc)

	scheduler := New(dbMock, mailer.New())
	err := scheduler.MergeAndSendEmails()
	if err != nil && err.Error() != "can not update status to valid" {
		t.Errorf("Error: %s", err)
	}
}
