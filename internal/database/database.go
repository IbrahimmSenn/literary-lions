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
	db, err := sql.Open("sqlite3", "/host/forum.db")
	if err != nil {
		return nil, err
	}

	// Test if the connection actually works
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//log.Println("Successfully connected to database!") <No more needed>
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

// CreateSessionsTable creates the Sessions table for user login sessions
func CreateSessionsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS sessions (
		id TEXT PRIMARY KEY,
		user_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		expires_at DATETIME NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("✅ Sessions table created successfully!")
	return nil
}

func CreateLikesTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS likes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER,
		comment_id INTEGER,
		is_like BOOLEAN NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (post_id) REFERENCES posts(id),
		FOREIGN KEY (comment_id) REFERENCES comments(id),
		UNIQUE(user_id, post_id),
		UNIQUE(user_id, comment_id)
	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("✅ Likes table created successfully!")
	return nil
}

// CreateCategoriesTable creates the Categories table
func CreateCategoriesTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL,
		slug TEXT UNIQUE NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("✅ Categories table created successfully!")
	return nil
}

// CreatePostCategoriesTable creates the junction table for post-category relationships
func CreatePostCategoriesTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS post_categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		category_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
		UNIQUE(post_id, category_id)
	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("✅ Post Categories table created successfully!")
	return nil
}

func InsertDefaultCategories(db *sql.DB) error {
	categories := []struct {
		name        string
		slug        string
		description string
	}{
		{"Book Reviews", "reviews", "Share your thoughts and reviews on books you've read"},
		{"Upcoming Releases", "upcoming", "Discuss upcoming book releases and new authors"},
		{"Chat with Authors", "chat-authors", "Connect and discuss with your favorite authors"},
		{"News and Interviews", "news", "Latest book industry news and author interviews"},
		{"Favourite Quotes", "quotes", "Share your favorite quotes from books"},
		{"Bookies Favourite 2025 Award", "25award", "Nominations and discussions for our annual award"},
		{"Book Recommendations", "recommendations", "Get and give book suggestions"},
		{"General Discussion", "general", "All other book-related conversations"},
	}

	for _, cat := range categories {
		query := `INSERT OR IGNORE INTO categories (name, slug, description) VALUES (?, ?, ?)`
		_, err := db.Exec(query, cat.name, cat.slug, cat.description)
		if err != nil {
			return err
		}
	}

	log.Println("✅ Default categories inserted successfully!")
	return nil
}
