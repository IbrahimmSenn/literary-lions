package handlers

import (
	"log"
	"net/http"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
	"literary-lions-forum/internal/utils"
)

// ForumHandler displays all posts with real-time vote counts from the database.
// Loads posts and their corresponding like/dislike counts for the forum page.(single user cannot like and dislike a post)
func ForumHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/forum" {
		http.NotFound(w, r)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	postRepo := sqlite.NewPostRepository(db)
	likeRepo := sqlite.NewLikeRepository(db)

	posts, err := postRepo.GetAllPosts()
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		http.Error(w, "Error loading posts", http.StatusInternalServerError)
		return
	}

	voteCounts := make(map[int]*models.VoteCount)
	for _, post := range posts {
		voteCount, err := likeRepo.GetPostVoteCount(post.ID)
		if err != nil {
			log.Printf("Error getting vote count for post %d: %v", post.ID, err)
			voteCount = &models.VoteCount{Likes: 0, Dislikes: 0}
		}
		voteCounts[post.ID] = voteCount
	}

	data := PageData{
		Result:     "",
		Status:     "200",
		Posts:      posts,
		VoteCounts: voteCounts,
	}

	err = Templates.ExecuteTemplate(w, "forum.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

// NewPostHandler manages post creation with authentication protection.
// Only authenticated users can access the form and create posts (LOGIN REQUIRED)
// Posts are automatically linked to the logged-in user's ID.
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/newpost" {
		http.NotFound(w, r)
		return
	}

	currentUser, err := utils.GetCurrentUser(r)
	if err != nil || currentUser == nil {
		log.Printf("🚫 Unauthorized attempt to access newpost")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

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

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		content := r.FormValue("content")

		if title == "" || content == "" {
			data := PageData{
				Error:  "Title and content are required",
				Status: "400",
			}
			Templates.ExecuteTemplate(w, "newpost.html", data)
			return
		}

		db, err := database.ConnectDB()
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		postRepo := sqlite.NewPostRepository(db)

		newPost := &models.Post{
			Title:   title,
			Content: content,
			UserID:  currentUser.ID,
		}

		err = postRepo.CreatePost(newPost)
		if err != nil {
			log.Printf("Error creating post: %v", err)
			http.Error(w, "Error saving post", http.StatusInternalServerError)
			return
		}

		log.Printf("✅ Post created by user %s: %s", currentUser.Username, title)
		http.Redirect(w, r, "/forum", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
