package repository

import "literary-lions-forum/internal/models"

// BookRepository defines operations for managing book recommendations
type BookRepository interface {
	// GetTodaysBookRecommendation gets the daily rotating book recommendation
	GetTodaysBookRecommendation() (*models.Book, error)
}
