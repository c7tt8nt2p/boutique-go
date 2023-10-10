package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/cart"
	"github.com/kx-boutique/ent/generated/user"
)

type CartEntity struct {
	Id     string
	UserId string
}

type CartRepo interface {
	Save(ctx context.Context, userId uuid.UUID) (uuid.UUID, error)
	FindIdByUserId(ctx context.Context, userId uuid.UUID) (uuid.UUID, error)
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

func (r *cartRepo) Save(ctx context.Context, userId uuid.UUID) (uuid.UUID, error) {
	entity, err := r.data.db.Cart.
		Create().
		SetUserID(userId).
		Save(ctx)

	if err != nil {
		return uuid.Nil, err
	}

	return entity.ID, nil
}

func (r *cartRepo) FindIdByUserId(ctx context.Context, userId uuid.UUID) (uuid.UUID, error) {
	entity, err := r.data.db.Cart.
		Query().
		Where(cart.HasUserIDOwnerWith(user.ID(userId))).
		Only(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return entity.ID, nil
}
