package util

import (
	"context"
	"fmt"
	ent "github.com/kx-boutique/ent/generated"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) (any, error)) (any, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	data, err := fn(tx)

	// rollback on error
	if err != nil {
		if rerr := rollback(tx); rerr != nil {
			return nil, rerr
		}
		return nil, err
	}

	// otherwise, commit
	if cerr := commit(tx); cerr != nil {
		return nil, cerr
	}

	return data, nil
}

func rollback(tx *ent.Tx) error {
	if err := tx.Rollback(); err != nil {
		err = fmt.Errorf("rolling back transaction: %w", err)
		return err
	}
	return nil
}

func commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		err = fmt.Errorf("committing transaction: %w", err)
		return err
	}
	return nil
}
