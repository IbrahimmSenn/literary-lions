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
		http.Redirect(w, r, "/login?message=login_required", http.StatusSeeOther)
		return
	}

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
	categoryRepo := sqlite.NewCategoryRepository(db)

	userPosts, err := postRepo.GetPostsByUserID(currentUser.ID)
	if err != nil {
		log.Printf("Error getting user posts: %v", err)
		userPosts = []*models.Post{}
	}

	userComments, err := commentRepo.GetCommentsByUserID(currentUser.ID)
	if err != nil {
		log.Printf("Error getting user comments: %v", err)
		userComments = []*models.Comment{}
	}

	//FOR THE LIKES
	likesGiven, err := likeRepo.GetLikesGivenByUserID(currentUser.ID)
	if err != nil {
		log.Printf("Error getting likes given: %v", err)
		likesGiven = 0
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
		LikesGiven:    likesGiven,
		LikesReceived: likesReceived,
	}

	// Get categories for user posts
	postCategories := make(map[int]*models.Category)
	for _, post := range userPosts {
		categories, err := categoryRepo.GetCategoriesForPost(post.ID)
		if err == nil && len(categories) > 0 {
			postCategories[post.ID] = categories[0]
		}
	}

	// Get vote counts for user posts
	voteCounts := make(map[int]*models.VoteCount)
	for _, post := range userPosts {
		voteCount, err := likeRepo.GetPostVoteCount(post.ID)
		if err == nil {
			voteCounts[post.ID] = voteCount
		}
	}

	// Get posts that user has liked
	likedPosts, err := likeRepo.GetPostsLikedByUser(currentUser.ID)
	if err != nil {
		log.Printf("Error getting liked posts: %v", err)
		likedPosts = []*models.Post{}
	}

	// Get categories for liked posts
	likedPostCategories := make(map[int]*models.Category)
	for _, post := range likedPosts {
		categories, err := categoryRepo.GetCategoriesForPost(post.ID)
		if err == nil && len(categories) > 0 {
			likedPostCategories[post.ID] = categories[0]
		}
	}

	// Get vote counts for liked posts
	likedPostVoteCounts := make(map[int]*models.VoteCount)
	for _, post := range likedPosts {
		voteCount, err := likeRepo.GetPostVoteCount(post.ID)
		if err == nil {
			likedPostVoteCounts[post.ID] = voteCount
		}
	}


	data := PageData{
		Result:              "",
		Status:              "200",
		CurrentUser:         currentUser,
		User:                currentUser,
		Stats:               stats,
		Posts:               userPosts,
		PostCategories:      postCategories,
		VoteCounts:          voteCounts,
		LikedPosts:          likedPosts,
		LikedPostCategories: likedPostCategories,
		LikedPostVoteCounts: likedPostVoteCounts,
		DarkMode:            GetDarkModePreference(r),

	}

	err = Templates.ExecuteTemplate(w, "user_profile.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}