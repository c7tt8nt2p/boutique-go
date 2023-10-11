package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/cartitem"
	"time"
)

type CartItemEntity struct {
	Id        uuid.UUID
	CartId    uuid.UUID
	ProductId uuid.UUID
	Qty       int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItemRepo interface {
	Save(ctx context.Context, cie *CartItemEntity) (*CartItemEntity, error)
	ExistsByCartIdAndProductId(ctx context.Context, cartId uuid.UUID, productId uuid.UUID) (bool, error)
}

type cartItemRepo struct {
	data *Data
	log  *log.Helper
}

func NewCartItemRepo(data *Data, logger log.Logger) CartItemRepo {
	return &cartItemRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/cart_item")),
	}
}

func (r *cartItemRepo) Save(ctx context.Context, cie *CartItemEntity) (*CartItemEntity, error) {
	fmt.Println("[insert] cartId", cie.CartId)
	fmt.Println("[insert] ProductId", cie.ProductId)
	fmt.Println("[insert] Qty", cie.Qty)
	entity, err := r.data.db.CartItem.
		Create().
		SetCartID(cie.CartId).
		SetProductID(cie.ProductId).
		SetQty(cie.Qty).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return &CartItemEntity{
		Id:        entity.ID,
		CartId:    entity.CartID,
		ProductId: entity.ProductID,
		Qty:       entity.Qty,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

func (r *cartItemRepo) ExistsByCartIdAndProductId(ctx context.Context, cartId uuid.UUID, productId uuid.UUID) (bool, error) {
	ok, err := r.data.db.CartItem.
		Query().
		Where(
			cartitem.And(
				cartitem.CartID(cartId),
				cartitem.ProductID(productId),
			),
		).
		Exist(ctx)

	if err != nil {
		return false, err
	}

	return ok, nil
}
