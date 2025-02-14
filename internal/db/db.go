package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/morshedulmunna/pxomart-api/internal/env"
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	duration := env.GetDuration(maxIdleTime, time.Duration(2*time.Second))
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db.PingContext(ctx)

	return db, nil
}
