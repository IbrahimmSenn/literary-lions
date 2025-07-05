package models

import "time"

type Like struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    *int      `json:"post_id,omitempty"`    // Pointer because it can be null
	CommentID *int      `json:"comment_id,omitempty"` // Pointer because it can be null
	IsLike    bool      `json:"is_like"`              // true = like, false = dislike
	CreatedAt time.Time `json:"created_at"`
}

// VoteCount represents the like/dislike counts for a post or comment
type VoteCount struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}
