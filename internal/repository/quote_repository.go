package repository

import "literary-lions-forum/internal/models"

// QuoteRepository defines operations for managing quotes
type QuoteRepository interface {
	// GetTodaysLiteratureQuote gets the daily rotating quote
	GetTodaysLiteratureQuote() (*models.Quote, error)
}
