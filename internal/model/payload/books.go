package payload

import "time"

type Book struct {
	ID              string    `json:"id"`
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

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBookPayload struct {
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
}

type EditBookPayload struct {
	ID              string    `json:"-"`
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
}
