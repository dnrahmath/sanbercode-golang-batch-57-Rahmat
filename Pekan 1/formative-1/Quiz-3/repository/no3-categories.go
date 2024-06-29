package repository

import (
	"Quiz-3/structs"
	"database/sql"
)

// Mendapatkan semua kategori dari database
func GetAllCategories(db *sql.DB) ([]structs.Category, error) {
	query := `SELECT id, name, created_at, updated_at FROM category`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// Menambahkan kategori ke database
func InsertCategory(db *sql.DB, category structs.Category) error {
	query := `INSERT INTO category (name, created_at, updated_at) VALUES ($1, $2, $3)`

	_, err := db.Exec(query, category.Name, category.CreatedAt, category.UpdatedAt)
	return err
}

// Memperbarui data kategori di database
func UpdateCategory(db *sql.DB, category structs.Category) error {
	query := `UPDATE category SET name = $1, created_at = $2, updated_at = $3 WHERE id = $4`

	_, err := db.Exec(query, category.Name, category.CreatedAt, category.UpdatedAt, category.ID)
	return err
}

// Menghapus kategori dari database
func DeleteCategory(db *sql.DB, category structs.Category) error {
	query := `DELETE FROM category WHERE id = $1`

	_, err := db.Exec(query, category.ID)
	return err
}
