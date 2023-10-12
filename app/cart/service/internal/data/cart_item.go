package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/cart/service/internal/biz"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/ent/generated/cartitem"
	entModel "github.com/kx-boutique/ent/model"
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

func (r *cartItemRepo) Save(ctx context.Context, client *ent.Client, cie *entModel.CartItem) *entModel.CartItem {
	ci, err := client.CartItem.
		Create().
		SetCartID(cie.CartId).
		SetProductID(cie.ProductId).
		SetQty(cie.Qty).
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

func (r *cartItemRepo) FindByCartId(ctx context.Context, client *ent.Client, cartId uuid.UUID) *entModel.CartWithProducts {
	ci, err := client.CartItem.
		Query().
		Where(
			cartitem.And(
				cartitem.CartID(cartId),
			),
		).
		Select(cartitem.FieldQty).
		QueryProductIDOwner().
		All(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	var cp []*entModel.CartProduct
	for _, e := range ci {
		cp = append(cp, &entModel.CartProduct{
			Id:    e.ID,
			Name:  e.Name,
			Price: e.UnitPrice,
		})
	}

	return &entModel.CartWithProducts{
		CartId:  cartId,
		Product: cp,
	}
}

func (r *cartItemRepo) ExistsByCartIdAndProductId(ctx context.Context, client *ent.Client, cartId uuid.UUID, productId uuid.UUID) bool {
	ok, err := client.CartItem.
		Query().
		Where(
			cartitem.And(
				cartitem.CartID(cartId),
				cartitem.ProductID(productId),
			),
		).
		Exist(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return ok
}
