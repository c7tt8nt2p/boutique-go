package data

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/kx-boutique/app/product/service/internal/conf"
	ent "github.com/kx-boutique/app/product/service/internal/data/ent/generated"
	_ "github.com/lib/pq"
)

var ProviderSet = wire.NewSet(NewEntClient, NewData, NewProductRepo)

type Data struct {
	db  *ent.Client
	log *log.Helper
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "product-service/data/ent"))

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
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "product-service/data"))

	d := &Data{
		db: entClient,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
