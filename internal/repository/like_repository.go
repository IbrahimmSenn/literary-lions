package repository

import "literary-lions-forum/internal/models"

// LikeRepository defines operations for managing likes and dislikes
type LikeRepository interface {
	AddLikeToPost(userID, postID int, isLike bool) error

	RemoveLikeFromPost(userID, postID int) error

	// GetPostVoteCount gets total likes and dislikes for a post
	GetPostVoteCount(postID int) (*models.VoteCount, error)

	// GetUserVoteOnPost checks if user has voted on a post and what type
	GetUserVoteOnPost(userID, postID int) (*models.Like, error)

	// AddLikeToComment adds a like/dislike to a comment
	AddLikeToComment(userID, commentID int, isLike bool) error

	// RemoveLikeFromComment removes a user's vote from a comment
	RemoveLikeFromComment(userID, commentID int) error

	// GetCommentVoteCount gets total likes and dislikes for a comment
	GetCommentVoteCount(commentID int) (*models.VoteCount, error)

	// GetUserVoteOnComment checks if user has voted on a comment and what type
	GetUserVoteOnComment(userID, commentID int) (*models.Like, error)

	// GetLikesGivenByUserID returns the count of likes given by a user
	GetLikesGivenByUserID(userID int) (int, error)

	GetPostsLikedByUser(userID int) ([]*models.Post, error)
}
