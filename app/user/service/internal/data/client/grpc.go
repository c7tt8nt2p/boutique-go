package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	authv1 "github.com/kx-boutique/api/auth/service/v1"
	cartv1 "github.com/kx-boutique/api/cart/service/v1"
	"github.com/kx-boutique/app/user/service/internal/conf"
)

func NewAuthClient(conf *conf.Client) (authv1.AuthClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(metadata.Client()),
		grpc.WithEndpoint(conf.AuthService.Url),
	)
	if err != nil {
		return nil, err
	}

	c := authv1.NewAuthClient(conn)

	return c, nil
}

func NewCartClient(conf *conf.Client) (cartv1.CartClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(metadata.Client()),
		grpc.WithEndpoint(conf.CartService.Url),
	)
	if err != nil {
		return nil, err
	}

	c := cartv1.NewCartClient(conn)

	return c, nil
}
