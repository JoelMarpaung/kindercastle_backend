package book

import (
	"context"
	"kindercastle_backend/internal/model/db"

	sq "github.com/Masterminds/squirrel"

	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/dbase"
)

func (r repository) EditByID(ctx context.Context, data payload.EditBookPayload, userID string) error {
	query, args, err := sq.
		Update(db.TableBook).
		SetMap(sq.Eq{
			"user_id":          userID,
			"title":            data.Title,
			"author":           data.Author,
			"isbn":             data.Isbn,
			"publisher":        data.Publisher,
			"publication_date": data.PublicationDate,
			"edition":          data.Edition,
			"genre":            data.Genre,
			"language":         data.Language,
			"number_of_pages":  data.NumberOfPages,
			"description":      data.Description,
			"price":            data.Price,
			"format":           data.Format,
			"image_url":        data.ImageUrl,
		}).
		Where("id = ? AND is_not_archived = ?", data.ID, true).
		ToSql()
	if err != nil {
		return err
	}

	tx := dbase.GetTrxFromContext(ctx, r.DB)
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
