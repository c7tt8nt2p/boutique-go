package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	pb "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/user/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
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
	return nil, nil

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

type CustomerClaims struct {
	UserId string `json:"user_id,omitempty"`
	jwtv4.RegisteredClaims
}

type authKey struct{}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	claims, ok := jwt.FromContext(ctx)
	if ok {
		fmt.Println("	ok1", claims)
		iss := claims.(jwtv4.MapClaims)["iss"].(string)
		user_id := claims.(jwtv4.MapClaims)["user_id"].(string)
		fmt.Println("		>>>>> iss", iss)
		fmt.Println("		>>>>> user_id", user_id)
		//if cc, ok := claims.(*CustomerClaims); ok {
		//	fmt.Println("	ok2")
		//	fmt.Println("	>>> cc", cc)
		//}
		//fmt.Println("	>>> userid", claims["user_id"])
	}

	user, err := s.uc.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserByIdResp{
		Id:   user.Id.String(),
		Name: user.Name,
	}, nil
}

func (s *UserService) GetIdByEmail(ctx context.Context, req *pb.GetIdByEmailReq) (*pb.GetIdByEmailResp, error) {
	id, err := s.uc.GetIdByEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.GetIdByEmailResp{
		Id: id.String(),
	}, nil
}
