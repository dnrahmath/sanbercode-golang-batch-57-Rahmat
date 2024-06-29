package repository

import (
	"Quiz-3/structs"
	"database/sql"
)

// Mendapatkan buku berdasarkan category_id dari database
func GetBooksByCategoryId(db *sql.DB, categoryId int) ([]structs.Book, error) {
	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id FROM book WHERE category_id = $1`

	rows, err := db.Query(query, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var book structs.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.UpdatedAt, &book.CategoryID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Mendapatkan semua buku dari database
func GetAllBooks(db *sql.DB) ([]structs.Book, error) {
	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id FROM book`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var book structs.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CreatedAt, &book.UpdatedAt, &book.CategoryID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Menambahkan buku ke database
func InsertBook(db *sql.DB, book structs.Book) error {
	query := `INSERT INTO book (title, description, image_url, release_year, price, total_page, thickness, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := db.Exec(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID)
	return err
}

// Memperbarui data buku di database
func UpdateBook(db *sql.DB, book structs.Book) error {
	query := `UPDATE book SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8 WHERE id = $9`

	_, err := db.Exec(query, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.ID)
	return err
}

// Menghapus buku dari database
func DeleteBook(db *sql.DB, book structs.Book) error {
	query := `DELETE FROM book WHERE id = $1`

	_, err := db.Exec(query, book.ID)
	return err
}
