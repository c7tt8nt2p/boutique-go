package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
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
	SaveNewCart(ctx context.Context, userId uuid.UUID) (uuid.UUID, error)
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

func (r *cartRepo) SaveNewCart(ctx context.Context, userId uuid.UUID) (uuid.UUID, error) {
	saved, err := r.data.db.UserCart.
		Create().
		SetUserID(userId).
		Save(ctx)

	if err != nil {
		return uuid.Nil, err
	}

	return saved.ID, nil
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
