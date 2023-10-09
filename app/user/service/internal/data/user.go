package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	ent "github.com/kx-boutique/ent/generated"
)

type UserEntity struct {
	Id   uuid.UUID
	Name string
}

type UserRepo interface {
	GetEntClient() *ent.Client
	SaveUser(ctx context.Context, client *ent.Client, u *UserEntity) (*UserEntity, error)
	FindById(ctx context.Context, client *ent.Client, id uuid.UUID) (*UserEntity, error)
	DeleteById(ctx context.Context, client *ent.Client, id uuid.UUID) error
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (r *userRepo) GetEntClient() *ent.Client {
	return r.data.db
}

func (r *userRepo) SaveUser(ctx context.Context, client *ent.Client, ue *UserEntity) (*UserEntity, error) {
	saved, err := client.User.
		Create().
		SetName(ue.Name).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return &UserEntity{
		Id:   saved.ID,
		Name: saved.Name,
	}, nil
}

func (r *userRepo) FindById(ctx context.Context, client *ent.Client, id uuid.UUID) (*UserEntity, error) {
	u, err := client.User.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return &UserEntity{
		Id:   u.ID,
		Name: u.Name,
	}, nil
}

func (r *userRepo) DeleteById(ctx context.Context, client *ent.Client, id uuid.UUID) error {
	err := client.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
