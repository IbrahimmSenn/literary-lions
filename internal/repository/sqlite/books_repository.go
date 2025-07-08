package sqlite

import (
	"literary-lions-forum/internal/data"
	"literary-lions-forum/internal/models"
	"time"
)

type SQLiteBookRepository struct {
	// No database needed - we'll use the data file
}

func NewBookRepository() *SQLiteBookRepository {
	return &SQLiteBookRepository{}
}

// GetTodaysBookRecommendation returns a book that rotates daily based on day of year
func (r *SQLiteBookRepository) GetTodaysBookRecommendation() (*models.Book, error) {
	// Get day of year (1-365/366) to rotate books
	dayOfYear := time.Now().YearDay()

	// Get total number of books from our data file
	totalBooks := len(data.BookRecommendations)
	if totalBooks == 0 {
		return nil, nil
	}

	// Calculate which book to show (rotating daily)
	bookIndex := (dayOfYear - 1) % totalBooks

	// Get the book from our data
	bookData := data.BookRecommendations[bookIndex]

	// Convert to our Book model
	book := &models.Book{
		Title:       bookData.Title,
		Author:      bookData.Author,
		Genre:       bookData.Genre,
		ImageURL:    bookData.ImageURL,
		Description: bookData.Description,
		CreatedAt:   time.Now(),
	}

	return book, nil
}
