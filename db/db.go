package db

import (
	"errors"
	"sync"

	"gitlab.com/seqone/mailtick/types"
)

// DB implements app.DB, as a very simple in-memory database.
type db struct {
	emails []types.Email
	mu     sync.Mutex
}

type DB interface {
	SaveEmail(*types.Email) error
	FindPendingEmails() ([]types.Email, error)
	PendingEmailsToError(error, string) error
	PendingEmailsToValid() error
}

// New initializes and returns a new DB.
func New() (DB, error) {
	return &db{}, nil
}

// SaveEmail insert e if e.ID==0, or else update it.
func (db *db) SaveEmail(e *types.Email) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	e.ID = len(db.emails) + 1
	e.Status = types.Pending

	db.emails = append(db.emails, *e)

	// Pas forcément besoin de retourner une erreur vue qu'on utilise une db en mémoire
	return nil
}

// FindPendingEmails returns the pending emails
func (db *db) FindPendingEmails() (emails []types.Email, err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for k, email := range db.emails {
		if email.Status != types.Pending {
			continue
		}

		emails = append(emails, email)
		db.emails[k].Status = types.Sent
	}

	if len(emails) == 0 {
		err = errors.New(types.ErrNoPendingEmails)
	}

	return
}

// Les functions PendingEmailsToError et PendingEmailsToValid étais dans l'optique d'un vrais DB
// comme ça on pourrais voir les mails qui on échoué et les relancer ou les analyser pour voir pourquoi ils ont échoué.

// PendingEmailsToError sets the error of all pending emails by recipient.
func (db *db) PendingEmailsToError(err error, recipient string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	for k, email := range db.emails {
		if email.Status != types.Pending && email.Recipient != recipient {
			continue
		}

		db.emails[k].Status = types.ErrSendEmail
		db.emails[k].Error = err.Error()
	}

	return nil
}

// PendingEmailsToValid sets the status of all pending emails to valid
func (db *db) PendingEmailsToValid() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	for k, email := range db.emails {
		if email.Status != types.Pending {
			continue
		}

		db.emails[k].Status = types.Valid
	}

	return nil
}
