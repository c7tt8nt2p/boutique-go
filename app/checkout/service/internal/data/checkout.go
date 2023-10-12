package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/checkout/service/internal/biz"
	entModel "github.com/kx-boutique/app/checkout/service/internal/biz/model"
	"github.com/kx-boutique/ent/generated/checkout"
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

func (r *checkoutRepo) Save(ctx context.Context, p *entModel.Checkout) *entModel.Checkout {
	saved, err := r.data.db.Checkout.
		Create().
		SetName(p.Name).
		SetDescription(p.Description).
		SetStock(p.Stock).
		SetUnitPrice(p.UnitPrice).
		Save(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return &entModel.Checkout{
		Id:          saved.ID,
		Name:        saved.Name,
		Description: saved.Description,
		Stock:       saved.Stock,
		UnitPrice:   saved.UnitPrice,
		CreatedAt:   saved.CreatedAt,
		UpdatedAt:   saved.UpdatedAt,
	}
}

func (r *checkoutRepo) FindById(ctx context.Context, id uuid.UUID) *entModel.Checkout {
	p, err := r.data.db.Checkout.Get(ctx, id)
	if err != nil {
		panic(errors.AppInternalErr("Checkout not found"))
	}

	return &entModel.Checkout{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Stock:       p.Stock,
		UnitPrice:   p.UnitPrice,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (r *checkoutRepo) IsPurchasable(ctx context.Context, id uuid.UUID, purchaseQty int32) bool {
	ok, err := r.data.db.Checkout.
		Query().
		Where(
			checkout.And(
				checkout.ID(id),
				checkout.StockGTE(purchaseQty),
			),
		).
		Exist(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}
	return ok
}
