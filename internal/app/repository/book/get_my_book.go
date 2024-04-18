package book

import (
	"context"
	"kindercastle_backend/internal/pkg/helper"

	sq "github.com/Masterminds/squirrel"

	"kindercastle_backend/internal/model/db"
	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/dbase"
)

func (r repository) GetMyBook(ctx context.Context, param payload.PagingAndFilterPayload, userID string) ([]db.Book, int64, error) {
	var (
		tx    = dbase.GetTrxFromContext(ctx, r.DB)
		items = make([]db.Book, 0)
		total int64
	)

	query, args, err := r.
		baseGetMyBookQuery(param, userID, helper.ImplodeStructTag(db.Book{}, "db")).
		OrderBy("created_at DESC").
		Limit(uint64(param.Limit)).
		Offset(uint64(param.GetOffset())).
		ToSql()
	if err != nil {
		return items, 0, err
	}

	if err = tx.SelectContext(ctx, &items, query, args...); err != nil {
		return items, 0, err
	}

	queryCount, args, err := r.baseGetMyBookQuery(param, userID, "COUNT(id) AS total").ToSql()
	if err != nil {
		return items, 0, err
	}

	if err = tx.GetContext(ctx, &total, queryCount, args...); err != nil {
		return items, 0, err
	}

	return items, total, nil
}

func (r repository) baseGetMyBookQuery(param payload.PagingAndFilterPayload, userID string, columns ...string) sq.SelectBuilder {
	builder := sq.Select(columns...).From(db.TableBook)

	if param.Search != "" {
		builder = builder.Where("title LIKE ?", "%"+param.Search+"%")
	}

	builder = builder.Where("is_not_archived = ?", true)

	builder = builder.Where("user_id = ?", userID)

	return builder
}
