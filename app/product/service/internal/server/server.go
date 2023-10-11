package server

import (
	"github.com/google/wire"
	"github.com/kx-boutique/app/product/service/internal/data/client"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, client.NewAuthClient)
