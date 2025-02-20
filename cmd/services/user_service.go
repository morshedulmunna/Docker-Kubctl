package services

import (
	"github.com/morshedulmunna/pxomart-api/cmd/models"
	"github.com/morshedulmunna/pxomart-api/internal/store"
)

// User represents a simple user model (replace with your actual model).

// UserService defines the business logic for user management.
type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, user *models.User) (*models.User, error)
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
func (u *userService) CreateUser(user *models.User) (*models.User, error) {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUser implements UserService.
func (u *userService) GetUser(id string) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(id string, user *models.User) (*models.User, error) {
	panic("unimplemented")
}
