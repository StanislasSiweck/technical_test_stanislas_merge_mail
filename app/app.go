package app

import (
	"context"
	"encoding/json"
	"gitlab.com/seqone/mailtick/types"
	"io"
	"net/http"

	"gitlab.com/seqone/mailtick/db"
)

// App holds an http.Server.
type App struct {
	srv *http.Server

	db db.DB
}

// New returns a new App.
func New(db db.DB) *App {
	return &App{db: db}
}

// Listen starts listening on addr.
func (a *App) Listen(addr string) error {
	mux := http.NewServeMux()
	a.srv = &http.Server{Addr: addr, Handler: mux}

	// Peux étre ajouter un middleware pour gérer l'accès à cette route (authentification)
	// Pas trop vue de consigne sur ce point
	mux.HandleFunc("/mail", handleError(a.handleMail))

	return a.srv.ListenAndServe()
}

// Shutdown graceful shutdowns the server.
func (a *App) Shutdown() {
	if a.srv != nil {
		a.srv.Shutdown(context.Background())
	}
}

func (a *App) handleMail(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodPost:

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			return errHTTP{code: http.StatusBadRequest, cause: err}
		}

		var email types.Email
		err = json.Unmarshal(bytes, &email)
		if err != nil {
			return errHTTP{code: http.StatusBadRequest, cause: err}
		}

		err = a.db.SaveEmail(&email)
		if err != nil {
			return errHTTP{code: http.StatusInternalServerError, cause: err}
		}

		return nil
	}
	return nil
}
