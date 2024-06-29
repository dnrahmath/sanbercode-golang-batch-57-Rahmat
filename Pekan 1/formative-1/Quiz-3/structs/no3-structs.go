package structs

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       string    `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoryID  int       `json:"category_id"`
}

type ReqBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       string `json:"price"`
	TotalPage   int    `json:"total_page"`
	CategoryID  int    `json:"category_id"`
}

func (reqBook ReqBook) Convert() Book {
	return Book{
		Title:       reqBook.Title,
		Description: reqBook.Description,
		ImageURL:    reqBook.ImageURL,
		ReleaseYear: reqBook.ReleaseYear,
		Price:       reqBook.Price,
		TotalPage:   reqBook.TotalPage,
		Thickness:   "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CategoryID:  reqBook.CategoryID,
	}
}

// {
// 	"id": 1,
// 	"name": "Category Name",
// 	"created_at": "2024-06-29T12:00:00Z",
// 	"updated_at": "2024-06-29T12:00:00Z"
// }

// {
// 	"title": "Book Title",
// 	"description": "Book Description",
// 	"image_url": "https://example.com/book-image.jpg",
// 	"release_year": 2024,
// 	"price": "19.99",
// 	"total_page": 300,
// 	"category_id": 1
// }
