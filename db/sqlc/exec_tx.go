package db

import (
	"context"
	"fmt"
)

// Notice that this func is not exported and thus starts with small letter
// Cuz we dont want other packages to meddle with this func
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		// if error then the transaction rolls back and db goes to previous state ie, like this transaction never happened
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
