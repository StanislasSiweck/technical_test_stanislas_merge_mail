// A package to send email to the app
// build and run ./mailpost after the app started
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/seqone/mailtick/types"
)

func main() {
	b, _ := json.Marshal(types.Email{
		Recipient: "bob",
		Subject:   "subject",
		Body:      "body",
	})

	resp, err := http.Post("http://localhost:8181/mail", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
