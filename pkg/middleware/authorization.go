package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	authv1 "github.com/kx-boutique/api/auth/service/v1"
	"github.com/kx-boutique/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

const authorizationKey string = "Authorization"

func JWTValidation(authClient authv1.AuthClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			defer func() {
				// do something after the handler finishes
			}()
			ctx = appendAuthorizationToClientCtx(ctx)
			_, err = authClient.Validate(ctx, &emptypb.Empty{})
			if err != nil {
				return nil, errors.ErrUInvalidCredentials
			}
			return handler(ctx, req)
		}
	}
}

func appendAuthorizationToClientCtx(ctx context.Context) context.Context {
	tr, ok := transport.FromServerContext(ctx)
	if ok {
		token := tr.RequestHeader().Get(authorizationKey)
		ctx = metadata.AppendToClientContext(ctx, authorizationKey, token)
	}
	return ctx
}
