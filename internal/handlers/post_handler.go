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
	categoryRepo := sqlite.NewCategoryRepository(db)

	currentUser, _ := utils.GetCurrentUser(r)

	posts, err := postRepo.GetAllPosts()
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		http.Error(w, "Error loading posts", http.StatusInternalServerError)
		return
	}

	voteCounts := make(map[int]*models.VoteCount)
	postCategories := make(map[int]*models.Category)

	for _, post := range posts {
		// Get vote counts
		voteCount, err := likeRepo.GetPostVoteCount(post.ID)
		if err != nil {
			log.Printf("Error getting vote count for post %d: %v", post.ID, err)
			voteCount = &models.VoteCount{Likes: 0, Dislikes: 0}
		}
		voteCounts[post.ID] = voteCount

		// Get category for each post - SAFER VERSION
		categories, err := categoryRepo.GetCategoriesForPost(post.ID)
		if err != nil || len(categories) == 0 {
			// Create a default General category object
			postCategories[post.ID] = &models.Category{
				ID:   0,
				Name: "General",
				Slug: "general",
			}
		} else {
			postCategories[post.ID] = categories[0]
		}
	}

	data := PageData{
		Result:         "",
		Status:         "200",
		Posts:          posts,
		VoteCounts:     voteCounts,
		PostCategories: postCategories,
		CurrentUser:    currentUser,
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
// NewPostHandler manages post creation with authentication protection.
// Only authenticated users can access the form and create posts (LOGIN REQUIRED)
// Posts are automatically linked to the logged-in user's ID.

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/newpost" {
		http.NotFound(w, r)
		return
	}

	// Check authentication first
	currentUser, err := utils.GetCurrentUser(r)
	if err != nil || currentUser == nil {
		// Redirect with helpful context
		http.Redirect(w, r, "/login?message=login_to_post", http.StatusSeeOther)
		return
	}

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get categories for dropdown
	categoryRepo := sqlite.NewCategoryRepository(db)
	categories, err := categoryRepo.GetAllCategories()
	if err != nil {
		log.Printf("Error getting categories: %v", err)
		categories = []*models.Category{} // Empty slice on error
	}

	if r.Method == http.MethodGet {
		log.Printf("✅ User %s accessing newpost form", currentUser.Username)
		data := PageData{
			Result:      "",
			Status:      "200",
			Categories:  categories,
			CurrentUser: currentUser, // IMPORTANT: Pass current user to template
		}

		err := Templates.ExecuteTemplate(w, "user_newpost.html", data)
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
		categoryIDStr := r.FormValue("category")
		log.Printf("🔍 DEBUG: Received category ID: '%s'", categoryIDStr)

		if title == "" || content == "" {
			data := PageData{
				Error:       "Title and content are required",
				Status:      "400",
				Categories:  categories,
				CurrentUser: currentUser, // IMPORTANT: Pass current user for error rendering
			}
			Templates.ExecuteTemplate(w, "user_newpost.html", data) // Use correct template
			return
		}

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

		// Associate category with post if selected
		if categoryIDStr != "" {
			categoryID, err := strconv.Atoi(categoryIDStr)
			if err == nil && categoryID > 0 {
				err = categoryRepo.AddCategoryToPost(newPost.ID, categoryID)
				if err != nil {
					log.Printf("Error associating category with post: %v", err)
					// Don't fail the post creation, just log the error
				}
			}
		}

		log.Printf("✅ Post created by user %s: %s", currentUser.Username, title)
		http.Redirect(w, r, "/forum", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// PostDetailHandler displays a single post with all its comments and vote counts
func PostDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Get post ID from URL - we'll need to parse /post/123
	postIDStr := r.URL.Path[len("/post/"):]
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
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

	// Get the post
	postRepo := sqlite.NewPostRepository(db)
	post, err := postRepo.GetPostByID(postID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Get the username of the post author
	userRepo := sqlite.NewUserRepository(db)
	postAuthor, err := userRepo.GetUserByID(post.UserID)
	if err != nil {
		log.Printf("Error getting post author: %v", err)
		// Set a default username if we can't find the author
		postAuthor = &models.User{Username: "Unknown User"}
	}

	// Get vote count for this post
	likeRepo := sqlite.NewLikeRepository(db)
	voteCount, err := likeRepo.GetPostVoteCount(postID)
	if err != nil {
		voteCount = &models.VoteCount{Likes: 0, Dislikes: 0}
	}

	// Get comments for this post
	commentRepo := sqlite.NewCommentRepository(db)
	comments, err := commentRepo.GetCommentsByPostID(postID)
	if err != nil {
		log.Printf("Error getting comments: %v", err)
		comments = []*models.Comment{} // Empty slice if error
	}

	// Get vote counts for each comment
	commentVoteCounts := make(map[int]*models.VoteCount)
	for _, comment := range comments {
		commentVoteCount, err := likeRepo.GetCommentVoteCount(comment.ID)
		if err != nil {
			log.Printf("Error getting vote count for comment %d: %v", comment.ID, err)
			commentVoteCount = &models.VoteCount{Likes: 0, Dislikes: 0}
		}
		commentVoteCounts[comment.ID] = commentVoteCount
	}

	currentUser, _ := utils.GetCurrentUser(r)

	data := PageData{
		Result:            "",
		Status:            "200",
		Post:              post,
		PostAuthor:        postAuthor, // Pass the author properly
		VoteCount:         voteCount,
		Comments:          comments,
		CurrentUser:       currentUser,
		CommentVoteCounts: commentVoteCounts,
	}

	err = Templates.ExecuteTemplate(w, "postview.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}
