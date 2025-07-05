package handlers

import (
	"log"
	"net/http"
	"strconv"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/repository/sqlite"
	"literary-lions-forum/internal/utils"
)

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	currentUser, err := utils.GetCurrentUser(r)
	if err != nil || currentUser == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	postIDStr := r.FormValue("post_id")
	action := r.FormValue("action") // "like", "dislike", or "remove"

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	likeRepo := sqlite.NewLikeRepository(db)

	// Handle the vote action
	switch action {
	case "like":
		err = likeRepo.AddLikeToPost(currentUser.ID, postID, true)
		log.Printf("👍 User %s liked post %d", currentUser.Username, postID)
	case "dislike":
		err = likeRepo.AddLikeToPost(currentUser.ID, postID, false)
		log.Printf("👎 User %s disliked post %d", currentUser.Username, postID)
	case "remove":
		err = likeRepo.RemoveLikeFromPost(currentUser.ID, postID)
		log.Printf("🚫 User %s removed vote from post %d", currentUser.Username, postID)
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Printf("Error handling vote: %v", err)
		http.Error(w, "Error processing vote", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}
