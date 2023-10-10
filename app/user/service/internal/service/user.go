package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/user/service/internal/biz"
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
	user, err := s.uc.RegisterNewUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResp{
		Id:    user.Id.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	user, err := s.uc.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserByIdResp{
		Id:   user.Id.String(),
		Name: user.Name,
	}, nil
}
