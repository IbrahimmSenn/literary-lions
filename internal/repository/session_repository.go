package repository

import "literary-lions-forum/internal/models"

// SessionRepository defines operations for managing user sessions
type SessionRepository interface {
	// CreateSession saves a new session to the database
	CreateSession(session *models.Session) error

	// GetSessionByID finds a session by its ID (for validating login)
	GetSessionByID(id string) (*models.Session, error)

	// DeleteSession removes a session (for logout)
	DeleteSession(id string) error

	// DeleteExpiredSessions removes all expired sessions (cleanup)
	DeleteExpiredSessions() error
}
