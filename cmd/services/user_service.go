package services

import "errors"

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
	users map[string]User // Simulating an in-memory DB
}

// NewUserService initializes and returns a UserService instance.
func NewUserService() UserService {
	return &userService{users: make(map[string]User)}
}

func (s *userService) CreateUser(user User) (User, error) {
	s.users[user.ID] = user
	return user, nil
}

func (s *userService) GetUser(id string) (User, error) {
	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) UpdateUser(id string, user User) (User, error) {
	_, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	s.users[id] = user
	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	_, exists := s.users[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}
