package db

import (
	"gitlab.com/seqone/mailtick/types"
	"sync"
)

type dbMock struct {
	emails []types.Email
	mu     sync.Mutex

	MockFunction
}

type MockFunction struct {
	SaveEmail            func(*types.Email) error
	FindPendingEmails    func() ([]types.Email, error)
	PendingEmailsToError func(error, string) error
	PendingEmailsToValid func() error
}

func NewMock(t MockFunction) (DB, error) {
	return &dbMock{
		MockFunction: t,
	}, nil
}

func (db *dbMock) SaveEmail(e *types.Email) error {
	return db.MockFunction.SaveEmail(e)
}

func (db *dbMock) FindPendingEmails() (emails []types.Email, err error) {
	return db.MockFunction.FindPendingEmails()
}

func (db *dbMock) PendingEmailsToError(err error, recipient string) error {
	return db.MockFunction.PendingEmailsToError(err, recipient)
}

func (db *dbMock) PendingEmailsToValid() error {
	return db.MockFunction.PendingEmailsToValid()
}
