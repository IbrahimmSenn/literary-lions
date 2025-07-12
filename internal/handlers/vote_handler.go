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

	// Get form values
	postIDStr := r.FormValue("post_id")
	commentIDStr := r.FormValue("comment_id")
	action := r.FormValue("action") // "like", "dislike", or "remove"

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	likeRepo := sqlite.NewLikeRepository(db)

	// Determine if this is a post vote or comment vote
	if commentIDStr != "" {
		// Comment vote
		commentID, err := strconv.Atoi(commentIDStr)
		if err != nil {
			http.Error(w, "Invalid comment ID", http.StatusBadRequest)
			return
		}

		// Check if user already has a vote on this comment
		existingVote, err := likeRepo.GetUserVoteOnComment(currentUser.ID, commentID)
		if err == nil && existingVote != nil {
			// User already has a vote on this comment
			if (action == "like" && existingVote.IsLike) || (action == "dislike" && !existingVote.IsLike) {
				// Same vote - remove it
				err = likeRepo.RemoveLikeFromComment(currentUser.ID, commentID)
				if err != nil {
					log.Printf("Error removing comment vote: %v", err)
					http.Error(w, "Error processing vote", http.StatusInternalServerError)
					return
				}
				log.Printf("🚫 User %s removed vote from comment %d", currentUser.Username, commentID)

				// Redirect
				redirectURL := r.FormValue("redirect")
				if redirectURL == "" {
					if postIDStr != "" {
						redirectURL = "/post/" + postIDStr
					} else {
						redirectURL = "/forum"
					}
				}
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)
				return
			}
		}

		// If we get here, either no existing vote or different vote - proceed with normal logic
		switch action {
		case "like":
			err = likeRepo.AddLikeToComment(currentUser.ID, commentID, true)
			log.Printf("👍 User %s liked comment %d", currentUser.Username, commentID)
		case "dislike":
			err = likeRepo.AddLikeToComment(currentUser.ID, commentID, false)
			log.Printf("👎 User %s disliked comment %d", currentUser.Username, commentID)
		case "remove":
			err = likeRepo.RemoveLikeFromComment(currentUser.ID, commentID)
			log.Printf("🚫 User %s removed vote from comment %d", currentUser.Username, commentID)
		default:
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		if err != nil {
			log.Printf("Error handling comment vote: %v", err)
			http.Error(w, "Error processing vote", http.StatusInternalServerError)
			return
		}

		// Get redirect parameter or default to post page
		redirectURL := r.FormValue("redirect")
		if redirectURL == "" {
			if postIDStr != "" {
				redirectURL = "/post/" + postIDStr
			} else {
				redirectURL = "/forum"
			}
		}
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)

	} else if postIDStr != "" {
		// Post vote
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		// Check if user already has a vote on this post
		existingVote, err := likeRepo.GetUserVoteOnPost(currentUser.ID, postID)
		if err == nil && existingVote != nil {
			// User already has a vote on this post
			if (action == "like" && existingVote.IsLike) || (action == "dislike" && !existingVote.IsLike) {
				// Same vote - remove it
				err = likeRepo.RemoveLikeFromPost(currentUser.ID, postID)
				if err != nil {
					log.Printf("Error removing post vote: %v", err)
					http.Error(w, "Error processing vote", http.StatusInternalServerError)
					return
				}
				log.Printf("🚫 User %s removed vote from post %d", currentUser.Username, postID)

				// Redirect
				redirectURL := r.FormValue("redirect")
				if redirectURL == "" {
					redirectURL = "/post/" + strconv.Itoa(postID)
				}
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)
				return
			}
		}

		// If we get here, either no existing vote or different vote - proceed with normal logic
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
			log.Printf("Error handling post vote: %v", err)
			http.Error(w, "Error processing vote", http.StatusInternalServerError)
			return
		}

		// Get redirect parameter or default to post page
		redirectURL := r.FormValue("redirect")
		if redirectURL == "" {
			redirectURL = "/post/" + strconv.Itoa(postID)
		}
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)

	} else {
		http.Error(w, "Missing post_id or comment_id", http.StatusBadRequest)
		return
	}
}
