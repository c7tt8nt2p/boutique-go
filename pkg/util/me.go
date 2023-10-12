package util

import (
	"context"
	"github.com/kx-boutique/pkg/model"
)

type MyselfKey struct{}

func Me(ctx context.Context) *model.Me {
	value := ctx.Value(MyselfKey{})
	return value.(*model.Me)
}
