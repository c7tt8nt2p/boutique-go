package server

import (
	"github.com/google/wire"
	"github.com/kx-boutique/app/checkout/service/internal/data/client"
)

var ProviderSet = wire.NewSet(NewGRPCServer, client.NewAuthClient)
