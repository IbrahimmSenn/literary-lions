package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/handlers"
	"literary-lions-forum/internal/middleware"
	"log"
	"net/http"
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

	// HELI ADDED THIS - This populates the categories table
	if err := database.InsertDefaultCategories(db); err != nil {
		log.Fatal("Failed to insert default categories:", err)
	}

	setupStaticFiles() // Move this BEFORE setupRoutes

	setupRoutes()

	startServer()
}

// Tables
func createTables(db *sql.DB) error {
	tables := []func(*sql.DB) error{
		database.CreateUsersTable,
		database.CreateCategoriesTable,
		database.CreatePostCategoriesTable,
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
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Page routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handle home page specifically
		if r.URL.Path == "/" {
			handlers.HomeHandler(w, r)
			return
		}
		// For any other unmatched path, show 404
		handlers.NotFoundHandler(w, r)
	})

	mux.HandleFunc("/error", handlers.ErrorHandler)
	mux.HandleFunc("/profile", handlers.ProfileHandler)

	// Authentication routes
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/logout", handlers.LogoutHandler)
	mux.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler)

	// Forum routes
	mux.HandleFunc("/forum", handlers.ForumHandler)
	mux.HandleFunc("/newpost", handlers.NewPostHandler)
	mux.HandleFunc("/post/", handlers.PostDetailHandler)

	mux.HandleFunc("/vote", handlers.VoteHandler)

	// Comment routes
	mux.HandleFunc("/comment", handlers.CommentHandler)

	// Light & Dark Mode
	mux.HandleFunc("/toggle-dark-mode", handlers.ToggleDarkModeHandler)

	// Wrap with recovery middleware and set as default handler
	handler := middleware.RecoveryMiddleware(mux)
	http.Handle("/", handler)
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
