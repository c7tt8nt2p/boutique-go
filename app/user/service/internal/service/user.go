package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/user/service/internal/biz"

	pb "github.com/kx-boutique/api/user/service/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user"))}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	user, err := s.uc.SaveUser(ctx, &biz.User{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResp{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
