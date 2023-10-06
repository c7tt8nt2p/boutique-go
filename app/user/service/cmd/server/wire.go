//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/kx-boutique/app/user/service/internal/biz"
	"github.com/kx-boutique/app/user/service/internal/client"
	"github.com/kx-boutique/app/user/service/internal/conf"
	"github.com/kx-boutique/app/user/service/internal/data"
	"github.com/kx-boutique/app/user/service/internal/server"
	"github.com/kx-boutique/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Client, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(client.ProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
