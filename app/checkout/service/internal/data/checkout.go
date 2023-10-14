package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/checkout/service/internal/biz"
	entModel "github.com/kx-boutique/app/checkout/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/pkg/errors"
)

type checkoutRepo struct {
	data *Data
	log  *log.Helper
}

func NewCheckoutRepo(data *Data, logger log.Logger) biz.CheckoutRepo {
	return &checkoutRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/checkout")),
	}
}

func (r *checkoutRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func (r *checkoutRepo) Save(ctx context.Context, client *ent.Client, m *entModel.Checkout) uuid.UUID {
	entity, err := client.Checkout.
		Create().
		SetUserID(m.UserId).
		SetTotalPrice(m.TotalPrice).
		Save(ctx)

	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return entity.ID
}
