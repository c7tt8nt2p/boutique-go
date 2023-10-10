package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/user/service/internal/conf"
	"github.com/kx-boutique/app/user/service/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})

	whiteList[v1.User_GetIdByEmail_FullMethodName] = struct{}{}

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger, s *service.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			validate.Validator(),
			selector.Server(
				jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
					return []byte("super-secret"), nil
				}),
			).Match(NewWhiteListMatcher()).Build(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, s)
	return srv
}
