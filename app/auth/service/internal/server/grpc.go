package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/kx-boutique/api/auth/service/v1"
	"github.com/kx-boutique/app/auth/service/internal/conf"
	"github.com/kx-boutique/app/auth/service/internal/service"
	server "github.com/kx-boutique/pkg/middleware"
)

var whitelist = map[string]struct{}{
	v1.Auth_Register_FullMethodName: {},
	v1.Auth_Login_FullMethodName:    {},
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(serverConf *conf.Server, appConf *conf.App, logger log.Logger, s *service.AuthService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
					return []byte(appConf.Jwt.Secret), nil
				}),
			).Match(server.NewWhiteListMatcher(whitelist)).Build(),
			logging.Server(logger),
			validate.Validator(),
		),
	}
	if serverConf.Grpc.Network != "" {
		opts = append(opts, grpc.Network(serverConf.Grpc.Network))
	}
	if serverConf.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(serverConf.Grpc.Addr))
	}
	if serverConf.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(serverConf.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterAuthServer(srv, s)
	return srv
}
