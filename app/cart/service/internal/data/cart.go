package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/kx-boutique/app/cart/service/internal/biz"
)

var _ biz.CartRepo = (*cartRepo)(nil)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/cart")),
	}
}

type Cart struct {
	UserId int64 `bson:"user_id"`
	Items  []struct {
		ItemId   int64 `bson:"item_id"`
		Quantity int64 `bson:"quantity"`
	} `bson:"items"`
}

func (r *cartRepo) GetCart(ctx context.Context, uid int64) (*biz.Cart, error) {
	po, err := r.data.db.Cart.Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	return &biz.Cart{
		UserId: po.ID,
	}, nil
}

func (r *cartRepo) DeleteCart(ctx context.Context, uid int64) error {
	//_, err := r.cartColl.DeleteOne(ctx, bson.M{"s": uid})
	return nil
}
func (r *cartRepo) SaveCart(ctx context.Context, c *biz.Cart) error {
	//items := bson.A{}
	//for _, x := range c.Items {
	//	items = append(items, bson.M{"item_id": x.Id, "quantity": x.Quantity})
	//}
	//result := r.cartColl.FindOneAndUpdate(ctx, bson.M{"s": c.UserId},
	//	bson.D{{"user_id", c.UserId}, {"items", items}})
	return nil
}
