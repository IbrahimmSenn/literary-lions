package models

import "time"

// User represents a user in our forum
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`              // I think we don't include password in JSON(security)
	CreatedAt time.Time `json:"created_at"`
} 