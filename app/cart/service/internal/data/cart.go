package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/cart/service/internal/biz"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/ent/generated/cart"
	"github.com/kx-boutique/ent/generated/user"
	"github.com/kx-boutique/pkg/errors"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func (r *cartRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/cart")),
	}
}

func (r *cartRepo) Save(ctx context.Context, client *ent.Client, userId uuid.UUID) uuid.UUID {
	entity, err := client.Cart.
		Create().
		SetUserID(userId).
		Save(ctx)

	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return entity.ID
}

func (r *cartRepo) FindIdByUserId(ctx context.Context, client *ent.Client, userId uuid.UUID) uuid.UUID {
	entity, err := client.Cart.
		Query().
		Where(cart.HasUserIDOwnerWith(user.ID(userId))).
		Only(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return entity.ID
}
