package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	authv1 "github.com/kx-boutique/api/auth/service/v1"
	"github.com/kx-boutique/pkg/errors"
	"github.com/kx-boutique/pkg/model"
	"github.com/kx-boutique/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

const authorizationKey string = "Authorization"

func JWTValidation(authClient authv1.AuthClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			defer func() {
			}()
			ctx = appendAuthorizationToClientCtx(ctx)

			validateResp, err := authClient.Validate(ctx, &emptypb.Empty{})
			if err != nil {
				return nil, errors.ErrUInvalidCredentials
			}

			ctx = appendMyselfToServerContext(ctx, validateResp)
			return handler(ctx, req)
		}
	}
}

// appendAuthorizationToClientCtx forwards token when calling to clients
func appendAuthorizationToClientCtx(ctx context.Context) context.Context {
	tr, ok := transport.FromServerContext(ctx)
	if ok {
		token := tr.RequestHeader().Get(authorizationKey)
		ctx = metadata.AppendToClientContext(ctx, authorizationKey, token)
	}
	return ctx
}

func appendMyselfToServerContext(ctx context.Context, validateResp *authv1.ValidateResp) context.Context {
	ctx = context.WithValue(ctx, util.MyselfKey{}, &model.Myself{
		UserId: validateResp.UserId,
		Email:  validateResp.Email,
	})
	return ctx
}
