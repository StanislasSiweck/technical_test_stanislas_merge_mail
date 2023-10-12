package db

import (
	"fmt"

	"gitlab.com/seqone/mailtick/types"
)

// DB implements app.DB, as a very simple in-memory database.
type db struct {
	emails []types.Email
}

type DB interface {
	SaveEmail(*types.Email) error
	FindPendingEmails() ([]types.Email, error)
}

// New initializes and returns a new DB.
func New() (DB, error) {
	return &db{}, nil
}

// SaveEmail insert e if e.ID==0, or else update it.
func (db *db) SaveEmail(e *types.Email) error {
	// TODO reminder to the concurrent accesses
	return fmt.Errorf("not implemented yet")
}

// FindPendingEmails returns the pending emails
func (db *db) FindPendingEmails() ([]types.Email, error) {
	// TODO reminder to the concurrent accesses
	return nil, fmt.Errorf("not implemented yet")
}
