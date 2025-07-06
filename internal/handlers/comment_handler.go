package handlers

import (
	"log"
	"net/http"
	"strconv"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
	"literary-lions-forum/internal/utils"
)

// CommentHandler handles comment creation on posts
func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if user is logged in
	currentUser, err := utils.GetCurrentUser(r)
	if err != nil || currentUser == nil {
		postIDStr := r.FormValue("post_id")
		// Redirect with a message parameter
		http.Redirect(w, r, "/login?message=login_to_comment&redirect=/post/"+postIDStr, http.StatusSeeOther)
		return
	}

	// Parse form data
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	postIDStr := r.FormValue("post_id")
	content := r.FormValue("comment")

	// Validate post ID
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Validate content
	if content == "" {
		http.Error(w, "Comment content cannot be empty", http.StatusBadRequest)
		return
	}

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create comment repository
	commentRepo := sqlite.NewCommentRepository(db)

	// Create new comment
	newComment := &models.Comment{
		Content: content,
		UserID:  currentUser.ID,
		PostID:  postID,
	}

	// Save comment to database
	err = commentRepo.CreateComment(newComment)
	if err != nil {
		log.Printf("Error creating comment: %v", err)
		http.Error(w, "Error saving comment", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Comment created by user %s on post %d", currentUser.Username, postID)

	// Redirect back to the post to see the new comment
	http.Redirect(w, r, "/post/"+strconv.Itoa(postID), http.StatusSeeOther)
}
