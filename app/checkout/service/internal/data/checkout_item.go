package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/checkout/service/internal/biz"
	entModel "github.com/kx-boutique/app/checkout/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/pkg/errors"
)

type checkoutItemRepo struct {
	data *Data
	log  *log.Helper
}

func NewCheckoutItemRepo(data *Data, logger log.Logger) biz.CheckoutItemRepo {
	return &checkoutItemRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/checkout_item")),
	}
}

func (r *checkoutItemRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func (r *checkoutItemRepo) SaveAll(ctx context.Context, client *ent.Client, cois []*entModel.CheckoutItem) []*ent.CheckoutItem {
	var coics []*ent.CheckoutItemCreate
	for _, coi := range cois {
		create := client.CheckoutItem.
			Create().
			SetCheckoutID(coi.CheckoutId).
			SetCartItemID(coi.CartItemId)
		coics = append(coics, create)
	}

	saved, err := client.CheckoutItem.CreateBulk(coics...).Save(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return saved
}
