package dbase

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"kindercastle_backend/internal/pkg/logging"
)

const trxKey = "db.transaction"

type (
	SQLExec interface {
		sqlx.Execer
		sqlx.ExecerContext
		NamedExec(query string, arg interface{}) (sql.Result, error)
		NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	}
	SQLQuery interface {
		sqlx.Queryer
		GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	}

	SQLQueryExec interface {
		SQLExec
		SQLQuery
	}

	WrapTransactionFunc func(ctx context.Context) error
)

func WithTransaction(ctx context.Context, db *sqlx.DB, fn WrapTransactionFunc, isolations ...sql.IsolationLevel) error {
	isolationLevel := sql.LevelRepeatableRead
	if len(isolations) >= 1 {
		isolationLevel = isolations[0]
	}

	tx, err := db.BeginTxx(ctx, &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			logging.Logger().Info().Msg("rollback on panic")
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
			logging.Logger().Info().Msg("rollback on error")
		} else {
			_ = tx.Commit()
			logging.Logger().Info().Msg("success committed")
		}
	}()

	// Passing DB transaction through context to repository.
	ctx = context.WithValue(ctx, trxKey, tx)

	err = fn(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetTrxFromContext(ctx context.Context, db *sqlx.DB) SQLQueryExec {
	tx, ok := ctx.Value(trxKey).(*sqlx.Tx)
	if ok {
		return tx
	}
	return db
}
