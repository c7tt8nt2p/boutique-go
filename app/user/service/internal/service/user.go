package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/user/service/internal/biz"

	v1 "github.com/kx-boutique/api/user/service/v1"
)

type UserService struct {
	v1.UnimplementedUserServer

	cc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(cc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		cc:  cc,
		log: log.NewHelper(log.With(logger, "module", "service/user"))}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserResp, error) {
	return &v1.CreateUserResp{}, nil
}
