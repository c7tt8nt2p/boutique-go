package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/auth/service/internal/biz"
	entModel "github.com/kx-boutique/app/auth/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/ent/generated/auth"
	"github.com/kx-boutique/pkg/errors"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/auth")),
	}
}

func (r *authRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func (r *authRepo) Save(ctx context.Context, client *ent.Client, ae *entModel.Auth) *entModel.Auth {
	entity, err := client.Auth.
		Create().
		SetPasswordHash(ae.PasswordHash).
		SetUserID(ae.UserId).
		Save(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return &entModel.Auth{
		Id:     entity.ID,
		UserId: entity.UserID,
	}
}

func (r *authRepo) FindPasswordHashByUserId(ctx context.Context, client *ent.Client, userId uuid.UUID) string {
	passwordHash, err := client.Auth.
		Query().
		Where(auth.UserID(userId)).
		Select(auth.FieldPasswordHash).
		String(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return passwordHash
}
func (r *authRepo) DeleteByUserId(ctx context.Context, client *ent.Client, userId uuid.UUID) uuid.UUID {
	_, err := client.Auth.
		Delete().
		Where(auth.UserID(userId)).
		Exec(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return userId
}
