package book

import (
	"context"
	"errors"
	"kindercastle_backend/internal/model/db"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/dbase"
)

func (r repository) Create(ctx context.Context, data payload.CreateBookPayload, userID string) (string, error) {
	genUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	id := genUUID.String()
	if id == "" {
		return "", errors.New("invalid parsing generated uuid")
	}

	bookMap := sq.Eq{
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
	}

	query, args, err := sq.
		Insert(db.TableBook).
		SetMap(bookMap).
		ToSql()
	if err != nil {
		return "", err
	}

	tx := dbase.GetTrxFromContext(ctx, r.DB)
	_, err = tx.ExecContext(ctx, query, args...)
	return id, err
}
