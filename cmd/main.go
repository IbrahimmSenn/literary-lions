package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
)

// Data structure for templates
type PageData struct {
	Input  string
	Result string
	Status string
	Error  string
	Posts  []*models.Post
}

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("web/templates/*.html"))
}

func main() {
	// Test our database connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close() // Close the database when the program ends

	// Create the Users table
	err = database.CreateUsersTable(db)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	// Create the Posts table
	err = database.CreatePostsTable(db)
	if err != nil {
		log.Fatal("Failed to create posts table:", err)
	}

	// Create the Comments table
	err = database.CreateCommentsTable(db)
	if err != nil {
		log.Fatal("Failed to create comments table:", err)
	}

	port := ":8080"

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/forum", forumHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/newpost", newpostHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/postview", postviewHandler)

	// Static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	// Start the server
	fmt.Printf("Server starting on port %s\n", port)
	log.Printf("Visit http://localhost%s in your browser\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Oi-oi, something went wrong", http.StatusMethodNotAllowed)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
	}

	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
	}

	err := templates.ExecuteTemplate(w, "login.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

func forumHandler(w http.ResponseWriter, r *http.Request) {
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
	// Add this debug line
	log.Printf("Found %d posts in database", len(posts))
	for i, post := range posts {
		log.Printf("Post %d: ID=%d, Title=%s", i, post.ID, post.Title)
	}

	data := PageData{
		Result: "",
		Status: "200",
		Posts:  posts,
	}

	err = templates.ExecuteTemplate(w, "forum.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
	}

	err := templates.ExecuteTemplate(w, "register.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

func newpostHandler(w http.ResponseWriter, r *http.Request) {
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

		err := templates.ExecuteTemplate(w, "newpost.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template error: %v", err)
		}
		return
	}

	// Handle POST request - process the form
	if r.Method == http.MethodPost {

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
			templates.ExecuteTemplate(w, "newpost.html", data)
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
			UserID:  1, // For now, I use fake user ID (we'll fix this when we add authentication)
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
func errorHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/error" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
	}

	err := templates.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

func postviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postview" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
	}

	err := templates.ExecuteTemplate(w, "postview.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}
