package data

import (
	"context"
	"github.com/google/uuid"

	"github.com/go-kratos/kratos/v2/log"
)

type UserEntity struct {
	Id   uuid.UUID
	Name string
}

type UserRepo interface {
	SaveUser(ctx context.Context, u *UserEntity) (*UserEntity, error)
	FindById(ctx context.Context, id uuid.UUID) (*UserEntity, error)
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

func (r *userRepo) SaveUser(ctx context.Context, ue *UserEntity) (*UserEntity, error) {
	saved, err := r.data.db.User.
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

func (r *userRepo) FindById(ctx context.Context, id uuid.UUID) (*UserEntity, error) {
	u, err := r.data.db.User.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return &UserEntity{
		Id:   u.ID,
		Name: u.Name,
	}, nil
}
