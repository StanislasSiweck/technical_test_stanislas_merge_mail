package mailer

import (
	"gitlab.com/seqone/mailtick/types"
)

// Mailer implements app.Mailer
type mockMailer struct {
	MockFunction
}

type MockFunction struct {
	Send func(e types.Email) error
}

func NewMock(m MockFunction) *mockMailer {
	return &mockMailer{
		MockFunction: m,
	}
}

// Send only prints e in stdout
func (m *mockMailer) Send(e types.Email) error {
	return m.MockFunction.Send(e)
}
