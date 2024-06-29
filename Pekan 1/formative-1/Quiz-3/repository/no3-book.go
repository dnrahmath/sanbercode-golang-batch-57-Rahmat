package repository

import (
	"Quiz-3/structs"
	"database/sql"
)

// Mendapatkan semua buku dari database
func GetAllBook(db *sql.DB) (results []structs.Book, err error) {
	sql := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id FROM Book`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book structs.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.UpdatedAt, &book.CategoryID)
		if err != nil {
			return nil, err
		}
		results = append(results, book)
	}

	return results, nil
}

// Menambahkan buku ke database
func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := `INSERT INTO Book (title, description, image_url, release_year, price, total_page, thickness, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = db.Exec(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID)
	return err
}

// Memperbarui data buku di database
func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := `UPDATE Book SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8 WHERE id = $9`

	_, err = db.Exec(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.ID)
	return err
}

// Menghapus buku dari database
func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := `DELETE FROM Book WHERE id = $1`

	_, err = db.Exec(sql, book.ID)
	return err
}
