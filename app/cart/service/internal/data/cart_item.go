package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/cart/service/internal/biz"
	entModel "github.com/kx-boutique/app/cart/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/ent/generated/cartitem"
	"github.com/kx-boutique/pkg/errors"
)

type cartItemRepo struct {
	data *Data
	log  *log.Helper
}

func NewCartItemRepo(data *Data, logger log.Logger) biz.CartItemRepo {
	return &cartItemRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/cart_item")),
	}
}

func (r *cartItemRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func (r *cartItemRepo) Save(ctx context.Context, client *ent.Client, m *entModel.CartItem) *entModel.CartItem {
	ci, err := client.CartItem.
		Create().
		SetCartID(m.CartId).
		SetProductID(m.ProductId).
		SetQty(m.Qty).
		Save(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return &entModel.CartItem{
		Id:        ci.ID,
		CartId:    ci.CartID,
		ProductId: ci.ProductID,
		Qty:       ci.Qty,
		CreatedAt: ci.CreatedAt,
		UpdatedAt: ci.UpdatedAt,
	}
}

func (r *cartItemRepo) FindByCartIdAndCheckedOutIsFalse(ctx context.Context, client *ent.Client, cartId uuid.UUID) *entModel.CartWithProducts {
	cis, err := client.Debug().CartItem.
		Query().
		WithProductIDOwner().
		Where(
			cartitem.And(
				cartitem.CartID(cartId),
				cartitem.CheckedOut(false),
			),
		).
		All(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	var cp []*entModel.CartProduct
	for _, ci := range cis {
		p := ci.Edges.ProductIDOwner
		cp = append(cp, &entModel.CartProduct{
			Id:        ci.ID,
			ProductId: ci.ProductID,
			Name:      p.Name,
			Price:     p.UnitPrice,
			Qty:       ci.Qty,
		})
	}

	return &entModel.CartWithProducts{
		CartId:  cartId,
		Product: cp,
	}
}

func (r *cartItemRepo) ExistsByCartIdAndProductIdAndCheckedOutIsFalse(ctx context.Context, client *ent.Client, cartId uuid.UUID, productId uuid.UUID) bool {
	ok, err := client.CartItem.
		Query().
		Where(
			cartitem.And(
				cartitem.CartID(cartId),
				cartitem.ProductID(productId),
				cartitem.CheckedOut(false),
			),
		).
		Exist(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return ok
}

func (r *cartItemRepo) UpdateCheckedOutTrueByIdsIn(ctx context.Context, client *ent.Client, ids []uuid.UUID) []uuid.UUID {
	_, err := client.CartItem.
		Update().
		SetCheckedOut(true).
		Where(
			cartitem.IDIn(ids...),
		).
		Save(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return ids
}

func (r *cartItemRepo) DeleteByCartIdAndProductIdAndCheckedOutIsFalse(ctx context.Context, client *ent.Client, cartId uuid.UUID, productId uuid.UUID) uuid.UUID {
	_, err := client.CartItem.
		Delete().
		Where(
			cartitem.And(
				cartitem.CartID(cartId),
				cartitem.ProductID(productId),
				cartitem.CheckedOut(false),
			),
		).Exec(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return productId
}
