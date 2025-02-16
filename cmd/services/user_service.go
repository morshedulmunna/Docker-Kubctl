package services

import (
	"github.com/morshedulmunna/pxomart-api/internal/store"
)

// User represents a simple user model (replace with your actual model).
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserService defines the business logic for user management.
type UserService interface {
	CreateUser(user User) (User, error)
	GetUser(id string) (User, error)
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	store store.UserStore
}

// NewUserService initializes and returns a UserService instance.
func NewUserService() UserService {
	return &userService{store: store.UserStore{}}
}

// CreateUser implements UserService.
func (u *userService) CreateUser(user User) (User, error) {
	// u.store.Create()
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUser implements UserService.
func (u *userService) GetUser(id string) (User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(id string, user User) (User, error) {
	panic("unimplemented")
}
