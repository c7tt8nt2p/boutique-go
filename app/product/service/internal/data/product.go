package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/product"
)

type ProductEntity struct {
	Id          string
	Name        string
	Description string
	Stock       int32
	UnitPrice   float64
}

type ProductRepo interface {
	Save(ctx context.Context, pe *ProductEntity) (*ProductEntity, error)
	FindById(ctx context.Context, id uuid.UUID) (*ProductEntity, error)
	IsPurchasable(ctx context.Context, id uuid.UUID, purchaseQty int32) (bool, error)
}

type productRepo struct {
	data *Data
	log  *log.Helper
}

func NewProductRepo(data *Data, logger log.Logger) ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/product")),
	}
}

func (r *productRepo) Save(ctx context.Context, pe *ProductEntity) (*ProductEntity, error) {
	saved, err := r.data.db.Product.
		Create().
		SetName(pe.Name).
		SetDescription(pe.Description).
		SetStock(pe.Stock).
		SetUnitPrice(pe.UnitPrice).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ProductEntity{
		Id:          saved.ID.String(),
		Name:        saved.Name,
		Description: saved.Description,
		Stock:       saved.Stock,
		UnitPrice:   saved.UnitPrice,
	}, nil
}

func (r *productRepo) FindById(ctx context.Context, id uuid.UUID) (*ProductEntity, error) {
	entity, err := r.data.db.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &ProductEntity{
		Id:          entity.ID.String(),
		Name:        entity.Name,
		Description: entity.Description,
		Stock:       entity.Stock,
		UnitPrice:   entity.UnitPrice,
	}, nil
}

func (r *productRepo) IsPurchasable(ctx context.Context, id uuid.UUID, purchaseQty int32) (bool, error) {
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
		return false, err
	}

	return ok, nil
}
