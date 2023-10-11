package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	authv1 "github.com/kx-boutique/api/auth/service/v1"
	"github.com/kx-boutique/api/product/service/v1"
	"github.com/kx-boutique/app/product/service/internal/conf"
	"github.com/kx-boutique/app/product/service/internal/service"
	"github.com/kx-boutique/pkg/middleware"
)

var whitelist = map[string]struct{}{}

func NewGRPCServer(c *conf.Server, logger log.Logger, authClient authv1.AuthClient, s *service.ProductService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			selector.Server(
				middleware.JWTValidation(authClient),
			).Match(middleware.NewWhiteListMatcher(whitelist)).Build(),
			logging.Server(logger),
			validate.Validator(),
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
	v1.RegisterProductServer(srv, s)
	return srv
}
