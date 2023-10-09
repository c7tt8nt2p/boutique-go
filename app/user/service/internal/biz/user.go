package biz

import (
	"context"
	v1 "github.com/kx-boutique/api/cart/service/v1"
	pb "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/user/service/internal/data"
	"github.com/kx-boutique/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type UserUseCase struct {
	cartClient v1.CartClient
	repo       data.UserRepo
	log        *log.Helper
}

func NewUserUseCase(cartClient v1.CartClient, repo data.UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		cartClient: cartClient,
		repo:       repo,
		log:        log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) SaveUser(ctx context.Context, req *pb.CreateUserReq) (*data.UserEntity, error) {
	user, err1 := uc.repo.SaveUser(ctx, uc.repo.GetEntClient(), &data.UserEntity{Name: req.Name})
	if err1 != nil {
		return nil, err1
	}

	_, err2 := uc.cartClient.NewCart(ctx, &v1.NewCartReq{
		UserId: user.Id.String(),
	})
	if err2 != nil {
		_ = uc.repo.DeleteById(ctx, uc.repo.GetEntClient(), user.Id)
		return nil, err2
	}

	return user, nil
}

func (uc *UserUseCase) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (*data.UserEntity, error) {
	client := uc.repo.GetEntClient()

	id, err := util.ParseUUID(req.Id)
	if err != nil {
		return nil, err
	}
	return uc.repo.FindById(ctx, client, id)
}

//func (uc *UserUseCase) GetUser(ctx context.Context, uid int64) (*User, error) {
//	return uc.repo.GetUser(ctx, uid)
//}
//
//func (uc *UserUseCase) DeleteUser(ctx context.Context, uid int64) error {
//	return uc.repo.DeleteUser(ctx, uid)
//}
