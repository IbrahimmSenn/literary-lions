package repository

import "literary-lions-forum/internal/models"

// CommentRepository defines operations we can do with comments in the database
type CommentRepository interface {
	// CreateComment adds a new comment to a post
	CreateComment(comment *models.Comment) error

	// GetCommentsByPostID finds all comments for a specific post
	GetCommentsByPostID(postID int) ([]*models.Comment, error)

	// GetCommentByID finds a specific comment by its ID
	GetCommentByID(id int) (*models.Comment, error)

	// GetCommentsByUserID finds all comments by a specific user
	GetCommentsByUserID(userID int) ([]*models.Comment, error)
}
