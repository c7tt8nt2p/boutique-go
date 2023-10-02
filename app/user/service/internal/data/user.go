package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/kx-boutique/app/user/service/internal/biz"
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

func (r *userRepo) SaveUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	saved, err := r.data.db.User.Create().SetName(u.Name).Save(ctx)

	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:   saved.ID.String(),
		Name: saved.Name,
	}, nil
}
