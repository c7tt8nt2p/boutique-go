package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	userv1 "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/auth/service/internal/conf"
)

func NewUserClient(conf *conf.Client) (userv1.UserClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(metadata.Client()),
		grpc.WithEndpoint(conf.UserService.Url),
	)
	if err != nil {
		return nil, err
	}

	c := userv1.NewUserClient(conn)

	return c, nil
}
