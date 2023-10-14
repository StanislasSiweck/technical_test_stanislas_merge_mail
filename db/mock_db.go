package db

import (
	"gitlab.com/seqone/mailtick/types"
	"sync"
)

type mockDb struct {
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
	return &mockDb{
		MockFunction: t,
	}, nil
}

func (db *mockDb) SaveEmail(e *types.Email) error {
	return db.MockFunction.SaveEmail(e)
}

func (db *mockDb) FindPendingEmails() (emails []types.Email, err error) {
	return db.MockFunction.FindPendingEmails()
}

func (db *mockDb) PendingEmailsToError(err error, recipient string) error {
	return db.MockFunction.PendingEmailsToError(err, recipient)
}

func (db *mockDb) PendingEmailsToValid() error {
	return db.MockFunction.PendingEmailsToValid()
}
