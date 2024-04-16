package book

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jinzhu/copier"

	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/logging"
)

func (s service) Detail(ctx context.Context, bookID string) (payload.Book, error) {
	var res payload.Book

	book, err := s.Book.GetByID(ctx, bookID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, custom.ErrInvalidBook
		}

		logging.Logger().Err(err).Any("bookID", bookID).Msg("failed to get book by detail")
		return res, custom.ErrInternalServer
	}

	err = copier.Copy(&res, &book)
	if err != nil {
		logging.Logger().Err(err).Msg("failed to transform response using copier.Copy")
		return res, custom.ErrInternalServer
	}

	return res, nil
}
