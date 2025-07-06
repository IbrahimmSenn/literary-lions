package sqlite

import (
	"database/sql"
	"literary-lions-forum/internal/models"
)

// SQLiteCommentRepository implements CommentRepository interface using SQLite
type SQLiteCommentRepository struct {
	db *sql.DB
}

// NewCommentRepository creates a new SQLite comment repository
func NewCommentRepository(db *sql.DB) *SQLiteCommentRepository {
	return &SQLiteCommentRepository{db: db}
}

// CreateComment adds a new comment to the database
func (r *SQLiteCommentRepository) CreateComment(comment *models.Comment) error {
	query := `
		INSERT INTO comments (content, user_id, post_id, created_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP) 
	`

	result, err := r.db.Exec(query, comment.Content, comment.UserID, comment.PostID)
	if err != nil {
		return err
	}

	// Get the ID of the newly created comment
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Update the comment struct with the new ID
	comment.ID = int(id)
	return nil
}

// GetCommentsByPostID finds all comments for a specific post with usernames
func (r *SQLiteCommentRepository) GetCommentsByPostID(postID int) ([]*models.Comment, error) {
	query := `
		SELECT 
			c.id, 
			c.content, 
			c.user_id, 
			c.post_id, 
			c.created_at,
			u.username
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ? 
		ORDER BY c.created_at ASC
	`

	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		comment := &models.Comment{}
		var username string
		err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID, &comment.CreatedAt, &username)
		if err != nil {
			return nil, err
		}

		// Add username to comment
		comment.Username = username
		comments = append(comments, comment)
	}

	return comments, nil
}

// GetCommentByID finds a comment by its ID
func (r *SQLiteCommentRepository) GetCommentByID(id int) (*models.Comment, error) {
	query := `SELECT id, content, user_id, post_id, created_at FROM comments WHERE id = ?`

	comment := &models.Comment{}
	err := r.db.QueryRow(query, id).Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentsByUserID finds all comments by a specific user
func (r *SQLiteCommentRepository) GetCommentsByUserID(userID int) ([]*models.Comment, error) {
	query := `SELECT id, content, user_id, post_id, created_at FROM comments WHERE user_id = ? ORDER BY created_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
