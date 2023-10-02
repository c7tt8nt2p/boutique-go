package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/user/service/internal/biz"

	pb "github.com/kx-boutique/api/user/service/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	cc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(cc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		cc:  cc,
		log: log.NewHelper(log.With(logger, "module", "service/user"))}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	return &pb.CreateUserResp{}, nil
}
