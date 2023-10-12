package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/kx-boutique/pkg/errors"
)

func AppRecovery() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			defer func() {
				if rerr := recover(); rerr != nil {
					appErr, ok := rerr.(*errors.AppErr)
					if ok {
						log.Context(ctx).Errorf("%v\n", appErr.Err)
						err = appErr.Err
					} else {
						// unknown errors are forwarded to Kratos recovery
						panic(rerr)
					}
				}
			}()
			return handler(ctx, req)
		}
	}
}
