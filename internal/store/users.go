package store

import (
	"database/sql"
	"errors"
)

type User struct {
	ID   string
	Name string
}

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

// Create inserts a new user into the database.
func (s *UserStore) Create(user User) (User, error) {
	if user.Name == "" {
		return User{}, errors.New("name is required")
	}

	query := "INSERT INTO users (name) VALUES (?) RETURNING id"
	var id string

	// Insert user and get the returned ID
	err := s.db.QueryRow(query, user.Name).Scan(&id)
	if err != nil {
		return User{}, err
	}

	user.ID = id
	return user, nil
}
