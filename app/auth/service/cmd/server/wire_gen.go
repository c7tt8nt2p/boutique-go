// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/auth/service/internal/biz"
	"github.com/kx-boutique/app/auth/service/internal/conf"
	"github.com/kx-boutique/app/auth/service/internal/data"
	"github.com/kx-boutique/app/auth/service/internal/server"
	"github.com/kx-boutique/app/auth/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewEntClient(confData, logger)
	dataData, cleanup, err := data.NewData(client, logger)
	if err != nil {
		return nil, nil, err
	}
	authRepo := data.NewAuthRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(authRepo, logger)
	authService := service.NewAuthService(authUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, authService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
