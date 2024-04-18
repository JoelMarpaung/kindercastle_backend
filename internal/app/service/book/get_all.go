package book

import (
	"context"
	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/logging"
)

func (s service) GetAll(ctx context.Context, param payload.PagingAndFilterPayload, userID string) ([]payload.Book, int64, error) {
	var items = make([]payload.Book, 0)

	books, count, err := s.Book.GetAll(ctx, param, userID)
	if err != nil {
		logging.Logger().Err(err).Msg("failed to get all paginated books")
		return items, 0, custom.ErrInternalServer
	}

	var result []payload.Book

	for _, data := range books {
		bookOwnership := data.UserID == userID
		res := payload.Book{
			ID:              data.ID,
			UserID:          data.UserID,
			Title:           data.Title,
			Author:          data.Author,
			Isbn:            data.Isbn,
			Publisher:       data.Publisher,
			PublicationDate: data.PublicationDate,
			Edition:         data.Edition,
			Genre:           data.Genre,
			Language:        data.Language,
			NumberOfPages:   data.NumberOfPages,
			Description:     data.Description,
			Price:           data.Price,
			Format:          data.Format,
			ImageUrl:        data.ImageUrl,
			BookOwnership:   bookOwnership,
			CreatedAt:       data.CreatedAt,
			UpdatedAt:       data.UpdatedAt,
		}

		result = append(result, res)
	}

	return result, count, nil
}
