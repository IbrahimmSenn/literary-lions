package sqlite

import (
	"database/sql"
	"literary-lions-forum/internal/models"
)

type SQLiteLikeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) *SQLiteLikeRepository {
	return &SQLiteLikeRepository{db: db}
}

func (r *SQLiteLikeRepository) AddLikeToPost(userID, postID int, isLike bool) error {
	query := `
		INSERT INTO likes (user_id, post_id, is_like, created_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(user_id, post_id) 
		DO UPDATE SET is_like = ?, created_at = CURRENT_TIMESTAMP
	`

	_, err := r.db.Exec(query, userID, postID, isLike, isLike)
	return err
}

func (r *SQLiteLikeRepository) RemoveLikeFromPost(userID, postID int) error {
	query := `DELETE FROM likes WHERE user_id = ? AND post_id = ?`
	_, err := r.db.Exec(query, userID, postID)
	return err
}

func (r *SQLiteLikeRepository) GetPostVoteCount(postID int) (*models.VoteCount, error) {
	query := `
		SELECT 
			COUNT(CASE WHEN is_like = 1 THEN 1 END) as likes,
			COUNT(CASE WHEN is_like = 0 THEN 1 END) as dislikes
		FROM likes 
		WHERE post_id = ?
	`

	voteCount := &models.VoteCount{}
	err := r.db.QueryRow(query, postID).Scan(&voteCount.Likes, &voteCount.Dislikes)
	if err != nil {
		return nil, err
	}

	return voteCount, nil
}

func (r *SQLiteLikeRepository) GetUserVoteOnPost(userID, postID int) (*models.Like, error) {
	query := `SELECT id, user_id, post_id, comment_id, is_like, created_at FROM likes WHERE user_id = ? AND post_id = ?`

	like := &models.Like{}
	err := r.db.QueryRow(query, userID, postID).Scan(&like.ID, &like.UserID, &like.PostID, &like.CommentID, &like.IsLike, &like.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No vote found
		}
		return nil, err
	}

	return like, nil
}

// GetUserVoteOnComment checks if user has voted on a comment and what type
func (r *SQLiteLikeRepository) GetUserVoteOnComment(userID, commentID int) (*models.Like, error) {
	query := `SELECT id, user_id, post_id, comment_id, is_like, created_at FROM likes WHERE user_id = ? AND comment_id = ?`

	like := &models.Like{}
	err := r.db.QueryRow(query, userID, commentID).Scan(&like.ID, &like.UserID, &like.PostID, &like.CommentID, &like.IsLike, &like.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No vote found
		}
		return nil, err
	}

	return like, nil
}

// AddLikeToComment adds or updates a like/dislike on a comment
func (r *SQLiteLikeRepository) AddLikeToComment(userID, commentID int, isLike bool) error {
	query := `
		INSERT INTO likes (user_id, comment_id, is_like, created_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(user_id, comment_id) 
		DO UPDATE SET is_like = ?, created_at = CURRENT_TIMESTAMP
	`

	_, err := r.db.Exec(query, userID, commentID, isLike, isLike)
	return err
}

func (r *SQLiteLikeRepository) RemoveLikeFromComment(userID, commentID int) error {
	query := `DELETE FROM likes WHERE user_id = ? AND comment_id = ?`
	_, err := r.db.Exec(query, userID, commentID)
	return err
}

// GetCommentVoteCount gets total likes and dislikes for a comment
func (r *SQLiteLikeRepository) GetCommentVoteCount(commentID int) (*models.VoteCount, error) {
	query := `
		SELECT 
			COUNT(CASE WHEN is_like = 1 THEN 1 END) as likes,
			COUNT(CASE WHEN is_like = 0 THEN 1 END) as dislikes
		FROM likes 
		WHERE comment_id = ?
	`

	voteCount := &models.VoteCount{}
	err := r.db.QueryRow(query, commentID).Scan(&voteCount.Likes, &voteCount.Dislikes)
	if err != nil {
		return nil, err
	}

	return voteCount, nil
}

// GetLikesGivenByUserID returns the count of likes given by a user, but only for posts not comments
func (r *SQLiteLikeRepository) GetLikesGivenByUserID(userID int) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM likes 
		WHERE user_id = ? AND is_like = 1 AND post_id IS NOT NULL
	`

	var count int
	err := r.db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetPostsLikedByUser gets all posts that a user has liked
func (r *SQLiteLikeRepository) GetPostsLikedByUser(userID int) ([]*models.Post, error) {
	query := `
		SELECT DISTINCT p.id, p.title, p.content, p.created_at, p.user_id
		FROM posts p
		JOIN likes l ON p.id = l.post_id
		WHERE l.user_id = ? AND l.is_like = 1
		ORDER BY l.created_at DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UserID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
