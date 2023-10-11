package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
)

func NewWhiteListMatcher(whitelist map[string]struct{}) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whitelist[operation]; ok {
			return false
		}
		return true
	}
}
