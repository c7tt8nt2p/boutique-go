package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CartEntity struct {
	UserId int64
	Items  []ItemEntity
}

type ItemEntity struct {
	Id       int64
	Quantity int64
}

type CartRepo interface {
	GetCart(ctx context.Context, id int64) (*CartEntity, error)
}

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func NewCartRepo(data *Data, logger log.Logger) CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/cart")),
	}
}

func (r *cartRepo) GetCart(ctx context.Context, uid int64) (*CartEntity, error) {
	//po, err := r.data.db.Cart.Get(ctx, uid)
	//if err != nil {
	//	return nil, err
	//}
	//return &CartEntity{
	//	UserId: po.ID,
	//}, nil
	return nil, nil
}
