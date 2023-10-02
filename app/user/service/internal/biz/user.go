package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	UserId int64
}

type UserRepo interface {
	GetUser(ctx context.Context, uid int64) (*User, error)
	SaveUser(ctx context.Context, c *User) error
	DeleteUser(ctx context.Context, uid int64) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) GetUser(ctx context.Context, uid int64) (*User, error) {
	return uc.repo.GetUser(ctx, uid)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, uid int64) error {
	return uc.repo.DeleteUser(ctx, uid)
}
