package util

import (
	"context"
	"github.com/kx-boutique/pkg/model"
)

type MyselfKey struct{}

func Myself(ctx context.Context) *model.Myself {
	value := ctx.Value(MyselfKey{})
	return value.(*model.Myself)
}
