package data

import (
	"context"
	"github.com/google/uuid"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductEntity struct {
	Id          string
	Name        string
	Description string
	Stock       int32
}

type ProductRepo interface {
	GetProductById(ctx context.Context, id string) (*ProductEntity, error)
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

func (r *productRepo) GetProductById(ctx context.Context, id string) (*ProductEntity, error) {
	buuid, err1 := uuid.FromBytes([]byte(id))
	if err1 != nil {
		return nil, err1
	}

	p, err2 := r.data.db.Product.Get(ctx, buuid)
	if err2 != nil {
		return nil, err2
	}

	return &ProductEntity{
		Id:          p.ID.String(),
		Name:        p.Name,
		Description: p.Description,
		Stock:       p.Stock,
	}, nil
}
