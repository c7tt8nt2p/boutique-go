package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/ent/generated/auth"
)

type AuthEntity struct {
	Id           uuid.UUID
	PasswordHash string
	UserId       uuid.UUID
}

type AuthRepo interface {
	GetEntClient() *ent.Client
	Save(ctx context.Context, client *ent.Client, ae *AuthEntity) (*AuthEntity, error)
	FindPasswordHashByUserId(ctx context.Context, client *ent.Client, userId uuid.UUID) (string, error)
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

func (r *authRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func (r *authRepo) Save(ctx context.Context, client *ent.Client, ae *AuthEntity) (*AuthEntity, error) {
	entity, err := client.Auth.
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

func (r *authRepo) FindPasswordHashByUserId(ctx context.Context, client *ent.Client, userId uuid.UUID) (string, error) {
	passwordHash, err := client.Auth.
		Query().
		Where(auth.UserID(userId)).
		Select(auth.FieldPasswordHash).
		String(ctx)

	fmt.Println("	passwordHash", passwordHash)

	if err != nil {
		return "", err
	}
	return passwordHash, err
}
