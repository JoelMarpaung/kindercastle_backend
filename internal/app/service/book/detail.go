package book

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path"

	"github.com/jinzhu/copier"

	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/logging"
)

func (s service) Detail(ctx context.Context, bookID string) (payload.Book, error) {
	var res payload.Book

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		logging.Logger().Error().Msg("BASE_URL is not set")
		return res, custom.ErrInternalServer
	}

	book, err := s.Book.GetByID(ctx, bookID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, custom.ErrInvalidBook
		}
		logging.Logger().Err(err).Str("bookID", bookID).Msg("failed to get book by detail")
		return res, custom.ErrInternalServer
	}

	err = copier.Copy(&res, &book)
	if err != nil {
		logging.Logger().Err(err).Msg("failed to transform response using copier.Copy")
		return res, custom.ErrInternalServer
	}

	res.ImageUrl = path.Join(baseURL, res.ImageUrl)

	return res, nil
}
