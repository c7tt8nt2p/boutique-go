package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type AuthEntity struct {
	Id           uuid.UUID
	PasswordHash string
	UserId       uuid.UUID
}

type AuthRepo interface {
	Save(ctx context.Context, ae *AuthEntity) (*AuthEntity, error)
}

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/product")),
	}
}

func (r *authRepo) Save(ctx context.Context, ae *AuthEntity) (*AuthEntity, error) {
	entity, err := r.data.db.Auth.
		Create().
		SetPasswordHash(ae.PasswordHash).
		SetUserID(ae.UserId).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &AuthEntity{
		Id:     entity.ID,
		UserId: entity.UserID,
	}, nil
}
