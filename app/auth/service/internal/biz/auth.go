package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	pb "github.com/kx-boutique/api/auth/service/v1"
	userv1 "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/auth/service/internal/biz/password"
	"github.com/kx-boutique/app/auth/service/internal/conf"
	"github.com/kx-boutique/app/auth/service/internal/data"
	"github.com/kx-boutique/pkg/model"
	"github.com/kx-boutique/pkg/util"
)

var (
	ErrInvalidCredentials = errors.Unauthorized("authentication failed", "invalid credentials")
)

type AuthUseCase struct {
	appConfig  *conf.App
	userClient userv1.UserClient
	repo       data.AuthRepo
	log        *log.Helper
}

func NewAuthUseCase(appConfig *conf.App, userClient userv1.UserClient, repo data.AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		appConfig:  appConfig,
		userClient: userClient,
		repo:       repo,
		log:        log.NewHelper(log.With(logger, "module", "usecase/auth"))}
}

func (uc *AuthUseCase) NewAuth(ctx context.Context, req *pb.RegisterReq) (*data.AuthEntity, error) {
	userId, err1 := util.ParseUUID(req.UserId)
	if err1 != nil {
		return nil, err1
	}

	hash, err2 := password.GeneratePasswordHash(req.Password)
	if err2 != nil {
		return nil, err2
	}

	entity, err3 := uc.repo.Save(ctx, uc.repo.GetEntClient(), &data.AuthEntity{
		UserId:       userId,
		PasswordHash: hash,
	})
	if err3 != nil {
		return nil, err3
	}

	return entity, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, req *pb.LoginReq) (string, error) {
	resp, err1 := uc.userClient.GetIdByEmail(ctx, &userv1.GetIdByEmailReq{
		Email: req.Email,
	})
	if err1 != nil {
		uc.log.Error(err1)
		return "", ErrInvalidCredentials
	}

	id, err2 := util.ParseUUID(resp.Id)
	if err2 != nil {
		uc.log.Error(err2)
		return "", ErrInvalidCredentials
	}

	passwordHash, err3 := uc.repo.FindPasswordHashByUserId(ctx, uc.repo.GetEntClient(), id)
	if err3 != nil {
		uc.log.Error(err3)
		return "", ErrInvalidCredentials
	}

	ok, _ := password.ComparePassword(req.Password, passwordHash)
	if !ok {
		return "", ErrInvalidCredentials
	}

	token, err4 := password.GenerateJWT(uc.appConfig.Jwt, id, req.Email)
	if err4 != nil {
		uc.log.Error(err4)
		return "", ErrInvalidCredentials
	}

	return token, nil
}

func (uc *AuthUseCase) ExtractJWTClaims(ctx context.Context) (*model.Myself, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		uc.log.Error("cannot extract claims from jwt")
		return nil, ErrInvalidCredentials
	}
	mapClaims := claims.(jwtv4.MapClaims)

	return &model.Myself{
		UserId: mapClaims["user_id"].(string),
		Email:  mapClaims["sub"].(string),
	}, nil
}
