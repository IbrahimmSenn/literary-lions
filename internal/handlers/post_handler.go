package handlers

import (
	"log"
	"net/http"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
	"literary-lions-forum/internal/utils"
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

// NewPostHandler handles post creation - PROTECTED ROUTE
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/newpost" {
		http.NotFound(w, r)
		return
	}

	// CHECK IF USER IS LOGGED IN - NEW SECURITY!
	currentUser, err := utils.GetCurrentUser(r)
	if err != nil || currentUser == nil {
		log.Printf("🚫 Unauthorized attempt to access newpost")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Handle GET request - show the form (only for logged-in users)
	if r.Method == http.MethodGet {
		log.Printf("✅ User %s accessing newpost form", currentUser.Username)
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

	// Handle POST request - process the form (only for logged-in users)
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

		// Create new post with REAL USER ID (not fake!)
		newPost := &models.Post{
			Title:   title,
			Content: content,
			UserID:  currentUser.ID, // Use real logged-in user!
		}

		// Save post to database
		err = postRepo.CreatePost(newPost)
		if err != nil {
			log.Printf("Error creating post: %v", err)
			http.Error(w, "Error saving post", http.StatusInternalServerError)
			return
		}

		log.Printf("✅ Post created by user %s: %s", currentUser.Username, title)

		// Redirect to forum page to see the new post
		http.Redirect(w, r, "/forum", http.StatusSeeOther)
		return
	}

	// Handle unsupported methods
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
