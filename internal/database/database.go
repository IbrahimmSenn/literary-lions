package database

import (
	"database/sql"
	"log"
	
	_ "github.com/mattn/go-sqlite3"
)

// ConnectDB creates a connection to our SQLite database
func ConnectDB() (*sql.DB, error) {
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

// CreateUsersTable creates the Users table
func CreateUsersTable(db *sql.DB) error {
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

// CreatePostsTable creates the Posts table
func CreatePostsTable(db *sql.DB) error {
	// SQL command to create the Posts table
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`
	
	// Execute the SQL command
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	
	log.Println("✅ Posts table created successfully!")
	return nil
}

// CreateCommentsTable creates the Comments table
func CreateCommentsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (post_id) REFERENCES posts(id)
	)`
	
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	
	log.Println("✅ Comments table created successfully!")
	return nil
}