package book

import (
	"context"
	"kindercastle_backend/internal/pkg/helper"

	sq "github.com/Masterminds/squirrel"

	"kindercastle_backend/internal/model/db"
	"kindercastle_backend/internal/pkg/dbase"
)

func (r repository) GetByID(ctx context.Context, bookID string) (db.Book, error) {
	var (
		res db.Book
		tx  = dbase.GetTrxFromContext(ctx, r.DB)
	)

	query, args, err := sq.
		Select(helper.ImplodeStructTag(db.Book{}, "db")).
		From(db.TableBook).
		Where("id = ? AND is_not_archived = ?", bookID, true).
		ToSql()
	if err != nil {
		return res, err
	}

	if err = tx.GetContext(ctx, &res, query, args...); err != nil {
		return res, err
	}

	return res, nil
}
