package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/user/service/internal/biz"
	entModel "github.com/kx-boutique/app/user/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/ent/generated/user"
	"github.com/kx-boutique/pkg/errors"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (r *userRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func (r *userRepo) SaveUser(ctx context.Context, client *ent.Client, u *entModel.User) *entModel.User {
	saved, err := client.User.
		Create().
		SetName(u.Name).
		SetEmail(u.Email).
		Save(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return &entModel.User{
		Id:        saved.ID,
		Name:      saved.Name,
		Email:     saved.Email,
		CreatedAt: saved.CreatedAt,
		UpdatedAt: saved.UpdatedAt,
	}
}

func (r *userRepo) FindById(ctx context.Context, client *ent.Client, id uuid.UUID) *entModel.User {
	u, err := client.User.Get(ctx, id)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return &entModel.User{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (r *userRepo) FindIdByEmail(ctx context.Context, client *ent.Client, email string) uuid.UUID {
	u, err := client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return u.ID
}

func (r *userRepo) DeleteById(ctx context.Context, client *ent.Client, id uuid.UUID) uuid.UUID {
	err := client.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	return id
}
