package utils

import (
	"net/http"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
)

// GetCurrentUser checks if user is logged in and returns the user
// Returns nil if not logged in or session expired
func GetCurrentUser(r *http.Request) (*models.User, error) {
	// Get session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, err // No cookie found
	}

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Check if session exists and is valid
	sessionRepo := sqlite.NewSessionRepository(db)
	session, err := sessionRepo.GetSessionByID(cookie.Value)
	if err != nil {
		return nil, err // Invalid or expired session
	}

	// Get user from session
	userRepo := sqlite.NewUserRepository(db)
	user, err := userRepo.GetUserByID(session.UserID)
	if err != nil {
		return nil, err // User not found
	}

	return user, nil
}

// IsLoggedIn checks if a user is currently logged in
func IsLoggedIn(r *http.Request) bool {
	user, err := GetCurrentUser(r)
	return err == nil && user != nil
}
