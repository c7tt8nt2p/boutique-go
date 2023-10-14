package util

import (
	"context"
	"fmt"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/pkg/errors"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) any) any {
	tx, err := client.Tx(ctx)
	if err != nil {
		panic(errors.AppInternalErr(fmt.Sprintf("Error initializing tx: %s", err.Error())))
	}

	defer func() {
		if v := recover(); v != nil {
			rerr := rollback(tx)
			if rerr != nil {
				panic(errors.AppInternalErr(rerr.Error()))
			}
			panic(v)
		}
	}()

	data := fn(tx)

	if cerr := commit(tx); cerr != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return data
}

func rollback(tx *ent.Tx) error {
	if err := tx.Rollback(); err != nil {
		err = fmt.Errorf("error rolling back transaction: %w", err)
		return err
	}
	return nil
}

func commit(tx *ent.Tx) error {
	if err := tx.Commit(); err != nil {
		err = fmt.Errorf("error committing transaction: %w", err)
		return err
	}
	return nil
}
