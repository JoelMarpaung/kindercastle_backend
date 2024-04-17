package db

import (
	"time"
)

var (
	TableBook = "books"
)

type Book struct {
	ID              string    `db:"id"`
	UserID          string    `db:"user_id"`
	Title           string    `db:"title"`
	Author          string    `db:"author"`
	Isbn            string    `db:"isbn"`
	Publisher       string    `db:"publisher"`
	PublicationDate time.Time `db:"publication_date"`
	Edition         string    `db:"edition"`
	Genre           string    `db:"genre"`
	Language        string    `db:"language"`
	NumberOfPages   int       `db:"number_of_pages"`
	Description     string    `db:"description"`
	Price           float64   `db:"price"`
	Format          string    `db:"format"`

	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
	DeletedAt     *time.Time `db:"deleted_at"`
	IsNotArchived bool       `db:"is_not_archived"` // auto generated column based on deletedAt null = true else false
}
