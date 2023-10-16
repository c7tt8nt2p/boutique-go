package data

import (
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/kx-boutique/app/checkout/service/internal/conf"
	"github.com/kx-boutique/app/checkout/service/internal/data/client"
	ent "github.com/kx-boutique/ent/generated"
	_ "github.com/lib/pq"
)

var ProviderSet = wire.NewSet(NewEntClient, NewData, NewCheckoutRepo, NewCheckoutItemRepo,
	client.NewAuthClient, client.NewCartClient, client.NewProductClient)

type Data struct {
	db  *ent.Client
	log *log.Helper
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "checkout-service/data/ent"))

	dataSource := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Url,
		conf.Database.DbName,
	)

	client, err := ent.Open(dialect.Postgres, dataSource)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	return client
}

func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "checkout-service/data"))

	d := &Data{
		db: entClient,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
