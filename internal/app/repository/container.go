package repository

import (
	"kindercastle_backend/internal/app/repository/book"
)

type Container struct {
	Book book.IRepository
}
