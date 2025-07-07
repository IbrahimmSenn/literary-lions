package models

import "time"

// Category represents a forum category
type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// PostCategory represents the relationship between posts and categories
type PostCategory struct {
	ID         int       `json:"id"`
	PostID     int       `json:"post_id"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

// PostWithCategory represents a post with its category information
type PostWithCategory struct {
	*Post
	CategoryName string `json:"category_name"`
	CategorySlug string `json:"category_slug"`
}
