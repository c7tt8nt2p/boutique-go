package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type ProductEntity struct {
	Id          string
	Name        string
	Description string
	Stock       int32
}

type ProductRepo interface {
	SaveProduct(ctx context.Context, pe *ProductEntity) (*ProductEntity, error)
	GetProductById(ctx context.Context, id uuid.UUID) (*ProductEntity, error)
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

func (r *productRepo) SaveProduct(ctx context.Context, pe *ProductEntity) (*ProductEntity, error) {
	saved, err := r.data.db.Product.
		Create().
		SetName(pe.Name).
		SetDescription(pe.Description).
		SetStock(pe.Stock).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ProductEntity{
		Id:          saved.ID.String(),
		Name:        saved.Name,
		Description: saved.Description,
		Stock:       saved.Stock,
	}, nil
}

func (r *productRepo) GetProductById(ctx context.Context, id uuid.UUID) (*ProductEntity, error) {
	p, err := r.data.db.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &ProductEntity{
		Id:          p.ID.String(),
		Name:        p.Name,
		Description: p.Description,
		Stock:       p.Stock,
	}, nil
}
