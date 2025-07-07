package sqlite

import (
	"database/sql"
	"literary-lions-forum/internal/models"
)

type SQLiteCategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *SQLiteCategoryRepository {
	return &SQLiteCategoryRepository{db: db}
}

// GetAllCategories retrieves all available categories
func (r *SQLiteCategoryRepository) GetAllCategories() ([]*models.Category, error) {
	query := `SELECT id, name, slug, description, created_at FROM categories ORDER BY name ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		category := &models.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Slug, &category.Description, &category.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// GetCategoryByID finds a category by its ID
func (r *SQLiteCategoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	query := `SELECT id, name, slug, description, created_at FROM categories WHERE id = ?`

	category := &models.Category{}
	err := r.db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.Slug, &category.Description, &category.CreatedAt)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategoryBySlug finds a category by its slug
func (r *SQLiteCategoryRepository) GetCategoryBySlug(slug string) (*models.Category, error) {
	query := `SELECT id, name, slug, description, created_at FROM categories WHERE slug = ?`

	category := &models.Category{}
	err := r.db.QueryRow(query, slug).Scan(&category.ID, &category.Name, &category.Slug, &category.Description, &category.CreatedAt)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// AddCategoryToPost associates a category with a post
func (r *SQLiteCategoryRepository) AddCategoryToPost(postID, categoryID int) error {
	query := `INSERT OR IGNORE INTO post_categories (post_id, category_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, postID, categoryID)
	return err
}

// GetCategoriesForPost gets all categories associated with a post
func (r *SQLiteCategoryRepository) GetCategoriesForPost(postID int) ([]*models.Category, error) {
	query := `
		SELECT c.id, c.name, c.slug, c.description, c.created_at
		FROM categories c
		JOIN post_categories pc ON c.id = pc.category_id
		WHERE pc.post_id = ?
		ORDER BY c.name ASC
	`

	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		category := &models.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Slug, &category.Description, &category.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// RemoveCategoriesFromPost removes all category associations from a post
func (r *SQLiteCategoryRepository) RemoveCategoriesFromPost(postID int) error {
	query := `DELETE FROM post_categories WHERE post_id = ?`
	_, err := r.db.Exec(query, postID)
	return err
}
