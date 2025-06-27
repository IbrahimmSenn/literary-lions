package main

import (
	"database/sql"
	"log"
	
	_ "github.com/mattn/go-sqlite3"
)

// This function creates a connection to our SQLite database
func connectDB() (*sql.DB, error) {
	// Open creates a database file called "forum.db" 
	// If it doesn't exist, SQLite will create it for us
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		return nil, err
	}
	
	// Test if the connection actually works
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	log.Println("Successfully connected to database!")
	return db, nil
}

// This function creates the Users table
func createUsersTable(db *sql.DB) error {
	// SQL command to create the Users table
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	
	// Execute the SQL command
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	
	log.Println("✅ Users table created successfully!")
	return nil
} 