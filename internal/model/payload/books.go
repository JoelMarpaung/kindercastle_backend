package payload

import "time"

type Book struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Isbn            string    `json:"isbn"`
	Publisher       string    `json:"publisher"`
	PublicationDate time.Time `json:"publication_date"`
	Edition         string    `json:"edition"`
	Genre           string    `json:"genre"`
	Language        string    `json:"language"`
	NumberOfPages   int       `json:"number_of_pages"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	Format          string    `json:"format"`
	ImageUrl        string    `json:"image_url"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBookPayload struct {
	UserID          string    `json:"user_id"`
	Title           string    `json:"title" validate:"required"`
	Author          string    `json:"author" validate:"required"`
	Isbn            string    `json:"isbn"`
	Publisher       string    `json:"publisher"`
	PublicationDate time.Time `json:"publication_date"`
	Edition         string    `json:"edition"`
	Genre           string    `json:"genre" validate:"required"`
	Language        string    `json:"language" validate:"required"`
	NumberOfPages   int       `json:"number_of_pages" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	Price           float64   `json:"price"`
	Format          string    `json:"format" validate:"required"`
	ImageUrl        string    `json:"image_url"`
}

type EditBookPayload struct {
	ID              string    `json:"-"`
	UserID          string    `json:"user_id"`
	Title           string    `json:"title" validate:"required"`
	Author          string    `json:"author" validate:"required"`
	Isbn            string    `json:"isbn"`
	Publisher       string    `json:"publisher"`
	PublicationDate time.Time `json:"publication_date"`
	Edition         string    `json:"edition"`
	Genre           string    `json:"genre" validate:"required"`
	Language        string    `json:"language" validate:"required"`
	NumberOfPages   int       `json:"number_of_pages" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	Price           float64   `json:"price"`
	Format          string    `json:"format" validate:"required"`
	ImageUrl        string    `json:"image_url"`
}
