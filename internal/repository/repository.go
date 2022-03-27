package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

type (
	txFunc func(ctx context.Context, tx *sql.Tx) error
)

func (r *repository) withTransaction(ctx context.Context, fn txFunc) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to prepare transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err = fn(ctx, tx); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return fmt.Errorf("error during transaction rollback cased by: %w, rollback failed due to: %s", err, txErr)
		}
		return fmt.Errorf("transaction rolled back due to: %w", err)
	}
	return tx.Commit()
}
