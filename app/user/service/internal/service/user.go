package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/user/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *UserService) GetMe(ctx context.Context, _ *emptypb.Empty) (*pb.GetMeResp, error) {
	u := s.uc.GetMe(ctx)

	return &pb.GetMeResp{
		Id:    u.Id.String(),
		Name:  u.Name,
		Email: u.Email,
	}, nil

}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	u := s.uc.RegisterNewUser(ctx, req)

	return &pb.RegisterResp{
		Id:        u.Id.String(),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	u := s.uc.GetUserById(ctx, req)

	return &pb.GetUserByIdResp{
		Id:        u.Id.String(),
		Name:      u.Name,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}, nil
}

func (s *UserService) GetIdByEmail(ctx context.Context, req *pb.GetIdByEmailReq) (*pb.GetIdByEmailResp, error) {
	id := s.uc.GetIdByEmail(ctx, req)

	return &pb.GetIdByEmailResp{
		Id: id.String(),
	}, nil
}
