package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/kx-boutique/api/cart/service/v1"
	"github.com/kx-boutique/app/user/service/internal/conf"
)

func NewCartClient(conf *conf.Client) (v1.CartClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(metadata.Client()),
		grpc.WithEndpoint(conf.CartService.Url),
	)
	if err != nil {
		return nil, err
	}

	c := v1.NewCartClient(conn)

	return c, nil
}
