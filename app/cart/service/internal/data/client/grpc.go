package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/kx-boutique/api/product/service/v1"
	"github.com/kx-boutique/app/cart/service/internal/conf"
)

func NewProductClient(conf *conf.Client) (v1.ProductClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(metadata.Client()),
		grpc.WithEndpoint(conf.ProductService.Url),
	)
	if err != nil {
		return nil, err
	}

	c := v1.NewProductClient(conn)

	return c, nil
}