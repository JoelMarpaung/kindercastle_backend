package book

import (
	"context"

	"github.com/jinzhu/copier"

	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/logging"
)

func (s service) GetAll(ctx context.Context, param payload.PagingAndFilterPayload) ([]payload.Book, int64, error) {
	var items = make([]payload.Book, 0)

	books, count, err := s.Book.GetAll(ctx, param)
	if err != nil {
		logging.Logger().Err(err).Msg("failed to get all paginated books")
		return items, 0, custom.ErrInternalServer
	}

	err = copier.Copy(&items, &books)
	if err != nil {
		logging.Logger().Err(err).Msg("failed to transform response using copier.Copy")
		return items, 0, custom.ErrInternalServer
	}

	return items, count, nil
}
