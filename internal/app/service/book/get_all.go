package book

import (
	"context"
	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/logging"
	"os"
	"path"
)

func (s service) GetAll(ctx context.Context, param payload.PagingAndFilterPayload, userID string) ([]payload.Book, int64, error) {
	var items = make([]payload.Book, 0)

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		logging.Logger().Error().Msg("BASE_URL is not set")
		return nil, 0, custom.ErrInternalServer
	}

	books, count, err := s.Book.GetAll(ctx, param, userID)
	if err != nil {
		logging.Logger().Err(err).Msg("failed to get all paginated books")
		return items, 0, custom.ErrInternalServer
	}

	var result []payload.Book

	for _, data := range books {
		bookOwnership := data.UserID == userID
		imageUrl := path.Join(baseURL, data.ImageUrl)
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
			ImageUrl:        imageUrl,
			BookOwnership:   bookOwnership,
			CreatedAt:       data.CreatedAt,
			UpdatedAt:       data.UpdatedAt,
		}

		result = append(result, res)
	}

	return result, count, nil
}
