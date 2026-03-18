package dbtools

import (
	"context"
	"errors"
	"github.com/fin/tools/pkg/helpers"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IQuery interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type TxRepository interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
	GetDb(ctx context.Context) IQuery
	GetNative() *pgxpool.Pool
}

type txKey string

const (
	TxCtxKey txKey = "tx"
)

func NewTxRepository(db *pgxpool.Pool) TxRepository {
	return &txRepository{
		db: db,
	}
}

type txRepository struct {
	db *pgxpool.Pool
}

type txState struct {
	tx pgx.Tx
	ok bool
}

func (t *txRepository) getTx(ctx context.Context) (pgx.Tx, bool) {
	state, ok := ctx.Value(TxCtxKey).(*txState)
	if !ok || state == nil {
		return nil, false
	}
	return state.tx, state.ok
}

func (t *txRepository) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	_, txExists := t.getTx(ctx)

	if txExists {
		return fn(ctx)
	}

	newTx, err := t.db.Begin(ctx)
	if err != nil {
		return helpers.Wrap("begin transaction failed", err)
	}

	ctx = context.WithValue(ctx, TxCtxKey, &txState{tx: newTx, ok: true})

	defer func() {
		if r := recover(); r != nil {
			_ = newTx.Rollback(context.Background())
			panic(r)
		}
	}()

	if err := fn(ctx); err != nil {
		if rbErr := newTx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
			return helpers.Wrap("rollback failed", rbErr)
		}
		return helpers.Wrap("transaction handler failed", err)
	}

	if err := newTx.Commit(ctx); err != nil {
		return helpers.Wrap("commit failed", err)
	}

	return nil
}

func (t *txRepository) GetDb(ctx context.Context) IQuery {
	tx, ok := t.getTx(ctx)
	if !ok {
		return t.db
	}
	return tx
}

func (t *txRepository) GetNative() *pgxpool.Pool {
	return t.db
}
