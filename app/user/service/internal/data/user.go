package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kx-boutique/app/cart/service/internal/biz"
)

var _ biz.CartRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(logger log.Logger) biz.CartRepo {
	return &userRepo{
		//data:     data,
		//cartColl: data.db.Collection("cart"),
		//log:      log.NewHelper(log.With(logger, "module", "repo/beer")),
	}
}

type Cart struct {
	UserId int64 `bson:"user_id"`
	Items  []struct {
		ItemId   int64 `bson:"item_id"`
		Quantity int64 `bson:"quantity"`
	} `bson:"items"`
}

func (r *userRepo) GetCart(ctx context.Context, uid int64) (*biz.Cart, error) {
	po, err := r.data.db.Cart.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Cart{
		UserId: po.ID,
	}, nil
}
