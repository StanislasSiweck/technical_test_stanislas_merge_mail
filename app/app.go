package app

import (
	"context"
	"fmt"
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
		// TODO get request body and transform into `types.Email`
		// Then call `a.db.SaveEmail()` to save the received Email
		return fmt.Errorf("not implemented")
	}
	return nil
}
