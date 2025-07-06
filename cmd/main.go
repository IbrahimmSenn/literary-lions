package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/handlers"
)

func init() {

	handlers.Templates = template.Must(template.ParseGlob("web/templates/*.html"))
}

func main() {
	// Database setup
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	setupRoutes()

	setupStaticFiles()

	startServer()
}

// Tables
func createTables(db *sql.DB) error {
	tables := []func(*sql.DB) error{
		database.CreateUsersTable,
		database.CreatePostsTable,
		database.CreateCommentsTable,
		database.CreateSessionsTable,
		database.CreateLikesTable,
	}

	for _, createTable := range tables {
		if err := createTable(db); err != nil {
			return err
		}
	}

	return nil
}

// Configuration
func setupRoutes() {
	// Page routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/error", handlers.ErrorHandler)
	http.HandleFunc("/postview", handlers.PostViewHandler)

	// Authentication routes
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	// Forum routes
	http.HandleFunc("/forum", handlers.ForumHandler)
	http.HandleFunc("/newpost", handlers.NewPostHandler)
	http.HandleFunc("/post/", handlers.PostDetailHandler)

	http.HandleFunc("/vote", handlers.VoteHandler)

	// Comment routes
	http.HandleFunc("/comment", handlers.CommentHandler)
}

// Static Configuration
func setupStaticFiles() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))
}

func startServer() {
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Printf("Visit http://localhost%s in your browser\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}
