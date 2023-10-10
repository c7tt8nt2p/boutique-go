package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/auth/service/v1"
	"github.com/kx-boutique/app/auth/service/internal/biz"
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

func (a *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	entity, err := a.uc.NewAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResp{
		Id: entity.Id.String(),
	}, nil
}

func (a *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return &pb.LoginResp{}, nil
}
