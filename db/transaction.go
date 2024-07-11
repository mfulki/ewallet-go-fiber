package db

import (
	"context"
	"database/sql"

	"github.com/mfulki/ewallet-go-fiber/constant"
	"github.com/sirupsen/logrus"
)

type Transaction interface {
	WithTransaction(ctx context.Context, tFunc func(context.Context) (any, error)) (any, error)
}

type dbTransaction struct {
	db *sql.DB
}

func NewDbTransaction(db *sql.DB) *dbTransaction {
	return &dbTransaction{
		db: db,
	}
}

func (t *dbTransaction) WithTransaction(ctx context.Context, tFunc func(context.Context) (any, error)) (any, error) {
	tx, err := t.db.Begin()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	txCtx := context.WithValue(ctx, constant.TxContext, tx)
	data, err := tFunc(txCtx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logrus.Error(err)
			return nil, err
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return data, nil
}
