package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
)

func MyMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			// do something before invoking the handler
			defer func() {
				// do something after the handler finishes
			}()
			return handler(ctx, req)
		}
	}
}
