package models

import "time"

// Quote represents a daily quote for the home page
type Quote struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Author    string    `json:"author"`
	BookTitle string    `json:"book_title"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
