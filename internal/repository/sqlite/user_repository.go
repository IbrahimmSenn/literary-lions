package sqlite

import (
	"database/sql"
	"literary-lions-forum/internal/models"
)

// SQLiteUserRepository implements UserRepository interface using SQLite
type SQLiteUserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new SQLite user repository
func NewUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

// CreateUser adds a new user to the database
func (r *SQLiteUserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, username, password, created_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP) 
	` //Questionmarks are for safely inserting the values into the database(prevents hacking)

	result, err := r.db.Exec(query, user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}

	// Get the ID of the newly created user
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Update the user struct with the new ID
	user.ID = int(id)
	return nil
}

// GetUserByEmail finds a user by their email - needed for login
func (r *SQLiteUserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, username, password, created_at FROM users WHERE email = ?`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err // User not found or database error
	}

	return user, nil
}

// GetUserByID finds a user by their ID - useful for checking sessions
func (r *SQLiteUserRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, email, username, password, created_at FROM users WHERE id = ?`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err // User not found or database error
	}

	return user, nil
}

// GetUserByUsername finds a user by their username - useful for checking duplicates
func (r *SQLiteUserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, email, username, password, created_at FROM users WHERE username = ?`

	user := &models.User{}
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err // User not found or database error
	}

	return user, nil
}
