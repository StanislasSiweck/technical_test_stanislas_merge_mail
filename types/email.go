package types

var Pending = "pending"
var Sent = "sent"
var Valid = "valid"
var ErrSendEmail = "error sending email"
var ErrNoPendingEmails = "no pending emails"

type Email struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Error     string `json:"error"`
}
