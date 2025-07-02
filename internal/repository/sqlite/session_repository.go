package sqlite

import (
	"database/sql"
	"literary-lions-forum/internal/models"
)

type SQLiteSessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SQLiteSessionRepository {
	return &SQLiteSessionRepository{db: db}
}

// CreateSession saves a new session to the database
func (r *SQLiteSessionRepository) CreateSession(session *models.Session) error {
	query := `
		INSERT INTO sessions (id, user_id, created_at, expires_at)
		VALUES (?, ?, CURRENT_TIMESTAMP, ?) 
	`

	_, err := r.db.Exec(query, session.ID, session.UserID, session.ExpiresAt)
	return err
}

// GetSessionByID finds a session by its ID and checks if it's still valid
func (r *SQLiteSessionRepository) GetSessionByID(id string) (*models.Session, error) {
	query := `
		SELECT id, user_id, created_at, expires_at 
		FROM sessions 
		WHERE id = ? AND expires_at > CURRENT_TIMESTAMP
	`

	session := &models.Session{}
	err := r.db.QueryRow(query, id).Scan(&session.ID, &session.UserID, &session.CreatedAt, &session.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// DeleteSession removes a session (for logout)
func (r *SQLiteSessionRepository) DeleteSession(id string) error {
	query := `DELETE FROM sessions WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// DeleteExpiredSessions removes all expired sessions (cleanup)
func (r *SQLiteSessionRepository) DeleteExpiredSessions() error {
	query := `DELETE FROM sessions WHERE expires_at <= CURRENT_TIMESTAMP`
	_, err := r.db.Exec(query)
	return err
}
