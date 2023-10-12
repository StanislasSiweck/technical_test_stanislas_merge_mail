package app

import (
	"fmt"
	"log"
	"net/http"
)

type errHTTP struct {
	code  int
	cause error
}

func (e errHTTP) Error() string {
	return fmt.Sprintf("%d %v", e.code, e.cause)
}

// handleError is a convenient middleware that permits to have handlers that
// returns an error. If the error is of type errHTTP, then errHTTP.code is used
// as the response status code, else code 500 is used.
func handleError(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("==> %s %s", r.Method, r.URL.Path)
		defer log.Printf("<== %s %s", r.Method, r.URL.Path)

		err := h(w, r)
		if err != nil {
			log.Printf("    %s %s error: %+v", r.Method, r.URL.Path, err)
			if errHTTP, ok := err.(errHTTP); ok {
				http.Error(w, errHTTP.cause.Error(), errHTTP.code)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
