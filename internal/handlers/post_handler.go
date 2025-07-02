package handlers

import (
	"log"
	"net/http"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
)

// ForumHandler displays all posts on the forum page
func ForumHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forum" {
		http.NotFound(w, r)
		return
	}

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create post repository
	postRepo := sqlite.NewPostRepository(db)

	// Get all posts from database
	posts, err := postRepo.GetAllPosts()
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		http.Error(w, "Error loading posts", http.StatusInternalServerError)
		return
	}

	// Debug logging
	log.Printf("Found %d posts in database", len(posts))
	for i, post := range posts {
		log.Printf("Post %d: ID=%d, Title=%s", i, post.ID, post.Title)
	}

	data := PageData{
		Result: "",
		Status: "200",
		Posts:  posts,
	}

	err = Templates.ExecuteTemplate(w, "forum.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

// NewPostHandler handles post creation
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/newpost" {
		http.NotFound(w, r)
		return
	}

	// Handle GET request - show the form
	if r.Method == http.MethodGet {
		data := PageData{
			Result: "",
			Status: "200",
		}

		err := Templates.ExecuteTemplate(w, "newpost.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template error: %v", err)
		}
		return
	}

	// Handle POST request - process the form
	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Get form values
		title := r.FormValue("title")
		content := r.FormValue("content")

		// Basic validation
		if title == "" || content == "" {
			data := PageData{
				Error:  "Title and content are required",
				Status: "400",
			}
			Templates.ExecuteTemplate(w, "newpost.html", data)
			return
		}

		// Connect to database
		db, err := database.ConnectDB()
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Create post repository
		postRepo := sqlite.NewPostRepository(db)

		// Create new post
		newPost := &models.Post{
			Title:   title,
			Content: content,
			UserID:  1, // For now, use fake user ID (we'll fix this when we add authentication)
		}

		// Save post to database
		err = postRepo.CreatePost(newPost)
		if err != nil {
			log.Printf("Error creating post: %v", err)
			http.Error(w, "Error saving post", http.StatusInternalServerError)
			return
		}

		// Redirect to forum page to see the new post
		http.Redirect(w, r, "/forum", http.StatusSeeOther)
		return
	}

	// Handle unsupported methods
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
