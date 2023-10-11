package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/auth/service/v1"
	"github.com/kx-boutique/app/auth/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	uc  *biz.AuthUseCase
	log *log.Helper
}

func NewAuthService(uc *biz.AuthUseCase, logger log.Logger) *AuthService {
	return &AuthService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/auth")),
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	entity, err := s.uc.NewAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResp{
		Id: entity.Id.String(),
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	accessToken, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResp{
		AccessToken: accessToken,
	}, nil
}

func (s *AuthService) Validate(ctx context.Context, req *emptypb.Empty) (*pb.ValidateResp, error) {
	data, err := s.uc.ExtractJWTClaims(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateResp{
		UserId: data.UserId,
		Email:  data.Email,
	}, nil
}
