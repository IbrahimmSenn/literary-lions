package repository

import "literary-lions-forum/internal/models"

// CommentRepository defines operations we can do with comments in the database
type CommentRepository interface {
	// CreateComment adds a new comment to a post
	CreateComment(comment *models.Comment) error

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
		
		// Add username to comment (we'll add this field to the model)
		comment.Username = username
		comments = append(comments, comment)
	}
	
	return comments, nil
}

	// GetCommentByID finds a specific comment by its ID
	GetCommentByID(id int) (*models.Comment, error)

	// GetCommentsByUserID finds all comments by a specific user
	GetCommentsByUserID(userID int) ([]*models.Comment, error)
}
