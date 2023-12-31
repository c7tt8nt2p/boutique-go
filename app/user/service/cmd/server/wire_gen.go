// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/user/service/internal/biz"
	"github.com/kx-boutique/app/user/service/internal/conf"
	"github.com/kx-boutique/app/user/service/internal/data"
	"github.com/kx-boutique/app/user/service/internal/data/client"
	"github.com/kx-boutique/app/user/service/internal/server"
	"github.com/kx-boutique/app/user/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confClient *conf.Client, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	authClient, err := client.NewAuthClient(confClient)
	if err != nil {
		return nil, nil, err
	}
	cartClient, err := client.NewCartClient(confClient)
	if err != nil {
		return nil, nil, err
	}
	generatedClient := data.NewEntClient(confData, logger)
	dataData, cleanup, err := data.NewData(generatedClient, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(cartClient, authClient, userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, authClient, userService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
