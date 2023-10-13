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

func (s *AuthService) NewAuth(ctx context.Context, req *pb.NewAuthReq) (*pb.NewAuthResp, error) {
	a := s.uc.NewAuth(ctx, req)

	return &pb.NewAuthResp{
		Id: a.Id.String(),
	}, nil
}

func (s *AuthService) DeleteAuth(ctx context.Context, req *pb.DeleteAuthReq) (*pb.DeleteAuthResp, error) {
	id := s.uc.DeleteAuth(ctx, req)

	return &pb.DeleteAuthResp{
		Id: id.String(),
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	accessToken := s.uc.Login(ctx, req)

	return &pb.LoginResp{
		AccessToken: accessToken,
	}, nil
}

func (s *AuthService) Validate(ctx context.Context, _ *emptypb.Empty) (*pb.ValidateResp, error) {
	me := s.uc.ExtractJWTClaims(ctx)

	return &pb.ValidateResp{
		UserId: me.UserId,
		Email:  me.Email,
	}, nil
}
