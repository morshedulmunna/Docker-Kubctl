package repositories

import (
	"database/sql"
	"errors"

	userModels "github.com/morshedulmunna/pxomart-api/app/models"
	config "github.com/morshedulmunna/pxomart-api/configs"
)

// AuthRepository interface
type AuthRepository interface {
	GetUserByEmail(email string) (*userModels.User, error)
}

// AuthRepositoryImpl struct
type AuthRepositoryImpl struct {
	db *sql.DB
}

// NewAuthRepository creates a new instance of AuthRepository
func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{db: config.DB}
}

// GetUserByEmail fetches user by email from PostgreSQL
func (repo *AuthRepositoryImpl) GetUserByEmail(email string) (*userModels.User, error) {
	var user userModels.User

	// Correct query to scan the proper columns into the user struct.
	query := "SELECT id::text, email, password FROM users WHERE email = $1"
	err := repo.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
