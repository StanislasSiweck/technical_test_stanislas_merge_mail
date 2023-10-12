package types

type Email struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Error     string `json:"error"`
}

type EmailRequest struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}
