package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
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
