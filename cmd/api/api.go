package main

import (
	"log"
	"net/http"
	"time"

	"github.com/morshedulmunna/pxomart-api/internal/store"
)

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr string
	db   dbConfig
}

// Application configuration struct
type application struct {
	config config
	store  store.Storage
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	defer srv.Close()

	log.Printf("Server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}
