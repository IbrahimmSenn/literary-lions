package sqlite

import (
	"literary-lions-forum/internal/data"
	"literary-lions-forum/internal/models"
	"time"
)

type SQLiteQuoteRepository struct {
	// No database needed - we'll use the data file
}

func NewQuoteRepository() *SQLiteQuoteRepository {
	return &SQLiteQuoteRepository{}
}

// GetTodaysLiteratureQuote returns a quote that rotates daily based on day of year
func (r *SQLiteQuoteRepository) GetTodaysLiteratureQuote() (*models.Quote, error) {
	// Get day of year (1-365/366) to rotate quotes
	dayOfYear := time.Now().YearDay()

	// Get total number of quotes from our data file
	totalQuotes := len(data.LiteratureQuotes)
	if totalQuotes == 0 {
		return nil, nil
	}

	// Calculate which quote to show (rotating daily)
	quoteIndex := (dayOfYear - 1) % totalQuotes

	// Get the quote from our data
	quoteData := data.LiteratureQuotes[quoteIndex]

	// Convert to our Quote model
	quote := &models.Quote{
		Text:      quoteData.Text,
		Author:    quoteData.Author,
		BookTitle: quoteData.BookTitle,
		CreatedAt: time.Now(),
	}

	return quote, nil
}
