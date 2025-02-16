package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

// Create inserts a new user into the database.

func (u *UserStore) Create(ctx context.Context) error {
	return nil
}
