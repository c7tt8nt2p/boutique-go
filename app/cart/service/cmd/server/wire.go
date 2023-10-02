//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/kx-boutique/app/cart/service/internal/biz"
	"github.com/kx-boutique/app/cart/service/internal/conf"
	"github.com/kx-boutique/app/cart/service/internal/data"
	"github.com/kx-boutique/app/cart/service/internal/server"
	"github.com/kx-boutique/app/cart/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
