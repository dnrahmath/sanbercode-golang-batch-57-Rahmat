package repository

import (
	"Quiz-3/structs"
	"database/sql"
)

// Mendapatkan semua kategori dari database
func GetAllCategories(db *sql.DB) (results []structs.Category, err error) {
	sql := `SELECT id, name, created_at, updated_at FROM Categories`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category structs.Category
		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, category)
	}

	return results, nil
}

// Menambahkan kategori ke database
func InsertCategories(db *sql.DB, category structs.Category) (err error) {
	sql := `INSERT INTO Categories (name, created_at, updated_at) VALUES ($1, $2, $3)`

	_, err = db.Exec(sql, category.Name, category.CreatedAt, category.UpdatedAt)
	return err
}

// Memperbarui data kategori di database
func UpdateCategories(db *sql.DB, category structs.Category) (err error) {
	sql := `UPDATE Categories SET name = $1, created_at = $2, updated_at = $3 WHERE id = $4`

	_, err = db.Exec(sql, category.Name, category.CreatedAt, category.UpdatedAt, category.ID)
	return err
}

// Menghapus kategori dari database
func DeleteCategories(db *sql.DB, category structs.Category) (err error) {
	sql := `DELETE FROM Categories WHERE id = $1`

	_, err = db.Exec(sql, category.ID)
	return err
}
