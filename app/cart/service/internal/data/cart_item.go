package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/ent/generated/cartitem"
	"time"
)

type CartItemProductEntity struct {
	Id          uuid.UUID
	Name        string
	Description string
	Stock       int32
	UnitPrice   float64
}

type CartItemEntity struct {
	Id        uuid.UUID
	CartId    uuid.UUID
	ProductId uuid.UUID
	Qty       int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItemRepo interface {
	GetEntClient() *ent.Client
	Save(ctx context.Context, client *ent.Client, cie *CartItemEntity) (*CartItemEntity, error)
	FindByCartId(ctx context.Context, client *ent.Client, cartId uuid.UUID) ([]*CartItemProductEntity, error)
	ExistsByCartIdAndProductId(ctx context.Context, client *ent.Client, cartId uuid.UUID, productId uuid.UUID) (bool, error)
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

func (r *cartItemRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func (r *cartItemRepo) Save(ctx context.Context, client *ent.Client, cie *CartItemEntity) (*CartItemEntity, error) {
	entity, err := client.CartItem.
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

func (r *cartItemRepo) FindByCartId(ctx context.Context, client *ent.Client, cartId uuid.UUID) ([]*CartItemProductEntity, error) {
	entity, err := client.CartItem.
		Query().
		Where(
			cartitem.And(
				cartitem.CartID(cartId),
			),
		).
		Select(cartitem.FieldQty).
		QueryProductIDOwner().
		All(ctx)
	var result []*CartItemProductEntity

	for _, e := range entity {
		result = append(result, &CartItemProductEntity{
			Id:          e.ID,
			Name:        e.Name,
			Description: e.Description,
			Stock:       e.Stock,
			UnitPrice:   e.UnitPrice,
		})
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *cartItemRepo) ExistsByCartIdAndProductId(ctx context.Context, client *ent.Client, cartId uuid.UUID, productId uuid.UUID) (bool, error) {
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
		return false, err
	}

	return ok, nil
}
