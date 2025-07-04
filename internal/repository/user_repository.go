package repository

import "literary-lions-forum/internal/models"

// UserRepository defines what operations we can do with users in the database
type UserRepository interface {
	// CreateUser adds a new user to the database
	CreateUser(user *models.User) error

	// GetUserByEmail finds a user by their email (for login)
	GetUserByEmail(email string) (*models.User, error)

	// GetUserByID finds a user by their ID (for sessions)
	GetUserByID(id int) (*models.User, error)

	GetUserByUsername(username string) (*models.User, error)
}
