package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/product/service/internal/biz"
	entModel "github.com/kx-boutique/app/product/service/internal/biz/model"
	"github.com/kx-boutique/ent/generated/product"
	"github.com/kx-boutique/pkg/errors"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/product")),
	}
}

func (r *productRepo) Save(ctx context.Context, p *entModel.Product) *entModel.Product {
	saved, err := r.data.db.Product.
		Create().
		SetName(p.Name).
		SetDescription(p.Description).
		SetStock(p.Stock).
		SetUnitPrice(p.UnitPrice).
		Save(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return &entModel.Product{
		Id:          saved.ID,
		Name:        saved.Name,
		Description: saved.Description,
		Stock:       saved.Stock,
		UnitPrice:   saved.UnitPrice,
		CreatedAt:   saved.CreatedAt,
		UpdatedAt:   saved.UpdatedAt,
	}
}

func (r *productRepo) FindById(ctx context.Context, id uuid.UUID) *entModel.Product {
	p, err := r.data.db.Product.Get(ctx, id)
	if err != nil {
		panic(errors.AppInternalErr("Product not found"))
	}

	return &entModel.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Stock:       p.Stock,
		UnitPrice:   p.UnitPrice,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (r *productRepo) IsPurchasable(ctx context.Context, id uuid.UUID, purchaseQty int32) bool {
	ok, err := r.data.db.Product.
		Query().
		Where(
			product.And(
				product.ID(id),
				product.StockGTE(purchaseQty),
			),
		).
		Exist(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}
	return ok
}
