package book

import (
	"context"
	"database/sql"
	"errors"
	"net/url"
	"os"

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

	base, err := url.Parse(baseURL)
	if err != nil {
		logging.Logger().Err(err).Msg("invalid base URL")
		return payload.Book{}, custom.ErrInternalServer
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

	imageEndpoint, _ := url.Parse(res.ImageUrl)
	imageUrl := base.ResolveReference(imageEndpoint).String()

	res.ImageUrl = imageUrl

	return res, nil
}
