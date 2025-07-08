package handlers

import (
	"log"
	"net/http"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
	"literary-lions-forum/internal/utils"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		http.NotFound(w, r)
		return
	}

	currentUser, err := utils.GetCurrentUser(r)
	if err != nil || currentUser == nil {
		// Redirect to login if not authenticated
		http.Redirect(w, r, "/login?message=login_required", http.StatusSeeOther)
		return
	}

	// Connect to database to get user stats
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Calculate user statistics
	postRepo := sqlite.NewPostRepository(db)
	commentRepo := sqlite.NewCommentRepository(db)
	likeRepo := sqlite.NewLikeRepository(db)

	userPosts, err := postRepo.GetPostsByUserID(currentUser.ID)
	if err != nil {
		log.Printf("Error getting user posts: %v", err)
		userPosts = []*models.Post{} // Empty slice on error
	}

	userComments, err := commentRepo.GetCommentsByUserID(currentUser.ID)
	if err != nil {
		log.Printf("Error getting user comments: %v", err)
		userComments = []*models.Comment{} // Empty slice on error
	}

	//FOR THE LIKES

	likesGiven, err := likeRepo.GetLikesGivenByUserID(currentUser.ID)
	if err != nil {
    	log.Printf("Error getting likes given: %v", err)
    	likesGiven = 0 // Default to 0 on error
	}

	likesReceived := 0
	for _, post := range userPosts {
		voteCount, err := likeRepo.GetPostVoteCount(post.ID)
		if err == nil {
			likesReceived += voteCount.Likes
		}
	}

	// Create user stats
	stats := &UserStats{
		PostCount:     len(userPosts),
		CommentCount:  len(userComments),
		LikesGiven: likesGiven,
		LikesReceived: likesReceived,

	}

	data := PageData{
		Result:      "",
		Status:      "200",
		CurrentUser: currentUser, // Add this for conditional headers
		User:        currentUser, // Keep this for backward compatibility
		Stats:       stats,
	}

	err = Templates.ExecuteTemplate(w, "user_profile.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}
