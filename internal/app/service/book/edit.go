package book

import (
	"context"
	"database/sql"
	"errors"
	"kindercastle_backend/internal/pkg/dbase"
	"kindercastle_backend/internal/pkg/logging"

	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/custom"
)

func (s service) Edit(ctx context.Context, data payload.EditBookPayload) error {
	_, err := s.Book.GetByID(ctx, data.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return custom.ErrInvalidBook
		}
		return custom.ErrInternalServer
	}

	err = dbase.WithTransaction(ctx, s.DB, func(ctx context.Context) error {
		err := s.Book.EditByID(ctx, data)

		return err
	})

	if err != nil {
		logging.Logger().Error().Err(err).Msg("failed to update book")
		return custom.ErrInternalServer
	}

	return nil
}
