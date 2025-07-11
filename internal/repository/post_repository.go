package repository

import "literary-lions-forum/internal/models"

// PostRepository defines all the  things we can do with posts in our database
type PostRepository interface {
	// CreatePost saves a brand new post when someone hits that "Post" button
	CreatePost(post *models.Post) error
	
	// GetPostByID finds one specific post - useful when someone clicks on a post
	GetPostByID(id int) (*models.Post, error)
	
	// GetAllPosts grabs every single post to show on the main forum page
	GetAllPosts() ([]*models.Post, error)
	
	// GetPostsByUserID finds all posts by one user ("show my posts" filter)
	GetPostsByUserID(userID int) ([]*models.Post, error)

	// GetPostsByCategory gets posts filtered by category
    GetPostsByCategory(category string) ([]*models.Post, error)
    
    // For sorting
    GetAllPostsSorted(sortBy string) ([]*models.Post, error)
    GetPostsByCategorySorted(category string, sortBy string) ([]*models.Post, error)
    
	// For search
	SearchPosts(query string, sortBy string) ([]*models.Post, error)
    SearchPostsByCategory(query string, category string, sortBy string) ([]*models.Post, error)
}
	
