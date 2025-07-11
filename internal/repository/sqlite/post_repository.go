package sqlite

import (
	"database/sql"
	"literary-lions-forum/internal/models"
)

// SQLitePostRepository implements PostRepository interface using SQLite
type SQLitePostRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new SQLite post repository
func NewPostRepository(db *sql.DB) *SQLitePostRepository {
	return &SQLitePostRepository{db: db}
}

// CreatePost adds a new post to the database
func (r *SQLitePostRepository) CreatePost(post *models.Post) error {
	query := `
		INSERT INTO posts (title, content, user_id, created_at, updated_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
	`
	
	result, err := r.db.Exec(query, post.Title, post.Content, post.UserID)
	if err != nil {
		return err
	}
	
	// Get the ID of the newly created post
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	// Update the post struct with the new ID
	post.ID = int(id)
	return nil
}

// GetPostByID finds a post by its ID
func (r *SQLitePostRepository) GetPostByID(id int) (*models.Post, error) {
	query := `SELECT id, title, content, user_id, created_at, updated_at FROM posts WHERE id = ?`
	
	post := &models.Post{}
	err := r.db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return post, nil
}

// GetAllPosts retrieves all posts ordered by creation date
func (r *SQLitePostRepository) GetAllPosts() ([]*models.Post, error) {
	query := `SELECT id, title, content, user_id, created_at, updated_at FROM posts ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var posts []*models.Post
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	
	return posts, nil
}

// GetPostsByUserID finds all posts by a specific user
func (r *SQLitePostRepository) GetPostsByUserID(userID int) ([]*models.Post, error) {
	query := `SELECT id, title, content, user_id, created_at, updated_at FROM posts WHERE user_id = ? ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var posts []*models.Post
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	
	return posts, nil
}

// GetPostsByUserIDWithLikes gets posts created by a user that have at least one like
func (r *SQLitePostRepository) GetPostsByUserIDWithLikes(userID int) ([]*models.Post, error) {
	query := `
        SELECT DISTINCT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at
        FROM posts p
        JOIN likes l ON p.id = l.post_id
        WHERE p.user_id = ? AND l.is_like = 1
        ORDER BY p.created_at DESC
    `
    
    rows, err := r.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var posts []*models.Post
    for rows.Next() {
        post := &models.Post{}
        err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    
    return posts, nil
}

// GetPostsByCategory gets posts filtered by category slug or name
func (r *SQLitePostRepository) GetPostsByCategory(category string) ([]*models.Post, error) {
    query := `
        SELECT DISTINCT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at
        FROM posts p
        JOIN post_categories pc ON p.id = pc.post_id
        JOIN categories c ON pc.category_id = c.id
        WHERE c.slug = ? OR c.name LIKE ?
        ORDER BY p.created_at DESC
    `
    
    categoryPattern := "%" + category + "%"
    
    rows, err := r.db.Query(query, category, categoryPattern)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var posts []*models.Post
    for rows.Next() {
        post := &models.Post{}
        err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    
    return posts, nil
}

// GetAllPostsSorted gets all posts with specified sorting
func (r *SQLitePostRepository) GetAllPostsSorted(sortBy string) ([]*models.Post, error) {
    var orderBy string
    
    switch sortBy {
    case "loved":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id AND is_like = 1
        ) DESC, p.created_at DESC`
    case "disliked":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id AND is_like = 0
        ) DESC, p.created_at DESC`
    case "recent":
        fallthrough
    default:
        orderBy = "ORDER BY p.created_at DESC"
    }
    
    query := `SELECT id, title, content, user_id, created_at, updated_at FROM posts p ` + orderBy
    
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var posts []*models.Post
    for rows.Next() {
        post := &models.Post{}
        err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    
    return posts, nil
}

// GetPostsByCategorySorted gets posts filtered by category with specified sorting
func (r *SQLitePostRepository) GetPostsByCategorySorted(category string, sortBy string) ([]*models.Post, error) {
    var orderBy string
    
    switch sortBy {
    case "liked":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id AND is_like = 1
        ) DESC, p.created_at DESC`
    case "disliked":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id AND is_like = 0
        ) DESC, p.created_at DESC`
    case "recent":
        fallthrough
    default:
        orderBy = "ORDER BY p.created_at DESC"
    }
    
    query := `
        SELECT DISTINCT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at
        FROM posts p
        JOIN post_categories pc ON p.id = pc.post_id
        JOIN categories c ON pc.category_id = c.id
        WHERE c.slug = ? OR c.name LIKE ?
    ` + orderBy
    
    categoryPattern := "%" + category + "%"
    
    rows, err := r.db.Query(query, category, categoryPattern)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var posts []*models.Post
    for rows.Next() {
        post := &models.Post{}
        err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    
    return posts, nil
}

// SearchPosts searches posts by title and content with sorting
func (r *SQLitePostRepository) SearchPosts(query string, sortBy string) ([]*models.Post, error) {
    var orderBy string
    
    switch sortBy {
    case "liked":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id AND is_like = 1
        ) DESC, p.created_at DESC`
    case "disliked":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id AND is_like = 0
        ) DESC, p.created_at DESC`
    case "viewed":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id
        ) DESC, p.created_at DESC`
    case "recent":
        fallthrough
    default:
        orderBy = "ORDER BY p.created_at DESC"
    }
    
    sqlQuery := `
        SELECT id, title, content, user_id, created_at, updated_at 
        FROM posts p 
        WHERE (title LIKE ? OR content LIKE ?)
    ` + orderBy
    
    searchPattern := "%" + query + "%"
    
    rows, err := r.db.Query(sqlQuery, searchPattern, searchPattern)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var posts []*models.Post
    for rows.Next() {
        post := &models.Post{}
        err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    
    return posts, nil
}

// SearchPostsByCategory searches posts by title and content within a specific category with sorting
func (r *SQLitePostRepository) SearchPostsByCategory(query string, category string, sortBy string) ([]*models.Post, error) {
    var orderBy string
    
    switch sortBy {
    case "liked":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id AND is_like = 1
        ) DESC, p.created_at DESC`
    case "disliked":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id AND is_like = 0
        ) DESC, p.created_at DESC`
    case "viewed":
        orderBy = `ORDER BY (
            SELECT COUNT(*) FROM likes WHERE post_id = p.id
        ) DESC, p.created_at DESC`
    case "recent":
        fallthrough
    default:
        orderBy = "ORDER BY p.created_at DESC"
    }
    
    sqlQuery := `
        SELECT DISTINCT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at
        FROM posts p
        JOIN post_categories pc ON p.id = pc.post_id
        JOIN categories c ON pc.category_id = c.id
        WHERE (p.title LIKE ? OR p.content LIKE ?) 
        AND (c.slug = ? OR c.name LIKE ?)
    ` + orderBy
    
    searchPattern := "%" + query + "%"
    categoryPattern := "%" + category + "%"
    
    rows, err := r.db.Query(sqlQuery, searchPattern, searchPattern, category, categoryPattern)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var posts []*models.Post
    for rows.Next() {
        post := &models.Post{}
        err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    
    return posts, nil
}