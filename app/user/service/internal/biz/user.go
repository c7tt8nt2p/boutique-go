package biz

import (
	"context"
	pb "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/user/service/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

type UserUseCase struct {
	repo data.UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo data.UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) SaveUser(ctx context.Context, req *pb.CreateUserReq) (*data.UserEntity, error) {
	return uc.repo.SaveUser(ctx, &data.UserEntity{Name: req.Name})
}

//func (uc *UserUseCase) GetUser(ctx context.Context, uid int64) (*User, error) {
//	return uc.repo.GetUser(ctx, uid)
//}
//
//func (uc *UserUseCase) DeleteUser(ctx context.Context, uid int64) error {
//	return uc.repo.DeleteUser(ctx, uid)
//}
