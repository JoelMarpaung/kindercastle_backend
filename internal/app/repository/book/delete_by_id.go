package book

import (
	"context"
	"kindercastle_backend/internal/model/db"

	sq "github.com/Masterminds/squirrel"

	"kindercastle_backend/internal/pkg/dbase"
)

func (r repository) DeleteByID(ctx context.Context, bookID string) error {
	tx := dbase.GetTrxFromContext(ctx, r.DB)

	query, args, err := sq.
		Update(db.TableBook).
		Set("deleted_at", sq.Expr("NOW()")).
		Where("id = ? AND is_not_archived = ?", bookID, true).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
