package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewCartRepo)

// Data .
type Data struct {
	db  *mongo.Database
	log *log.Helper
}
