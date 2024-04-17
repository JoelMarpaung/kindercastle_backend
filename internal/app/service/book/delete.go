package book

import (
	"context"
	"database/sql"
	"errors"

	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/dbase"
)

func (s service) Delete(ctx context.Context, bookID string, userID string) error {
	book, err := s.Book.GetByID(ctx, bookID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return custom.ErrInvalidBook
	}
	if err != nil {
		return custom.ErrInternalServer
	}

	return dbase.WithTransaction(ctx, s.DB, func(ctx context.Context) error {
		return s.Book.DeleteByID(ctx, book.ID, userID)
	})
}
