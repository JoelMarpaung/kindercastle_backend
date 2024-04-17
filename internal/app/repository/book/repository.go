package book

import (
	"context"

	"kindercastle_backend/internal/model/db"
	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg"
)

type IRepository interface {
	Create(ctx context.Context, data payload.CreateBookPayload, userID string) (string, error)
	EditByID(ctx context.Context, data payload.EditBookPayload, userID string) error
	GetByID(ctx context.Context, bookID string) (db.Book, error)
	GetAll(ctx context.Context, param payload.PagingAndFilterPayload) ([]db.Book, int64, error)
	DeleteByID(ctx context.Context, bookID string, userID string) error
}

type repository struct {
	*pkg.Options
}

func New(options *pkg.Options) IRepository {
	return repository{Options: options}
}
