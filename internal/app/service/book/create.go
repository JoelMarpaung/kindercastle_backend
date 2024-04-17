package book

import (
	"context"
	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/dbase"
	"kindercastle_backend/internal/pkg/logging"

	"kindercastle_backend/internal/model/payload"
)

func (s service) Create(ctx context.Context, data payload.CreateBookPayload, userID string) error {
	err := dbase.WithTransaction(ctx, s.DB, func(ctx context.Context) error {
		_, err := s.Book.Create(ctx, data, userID)

		return err
	})

	if err != nil {
		logging.Logger().Error().Err(err).Msg("failed to insert book")
		return custom.ErrInternalServer
	}

	return nil
}
