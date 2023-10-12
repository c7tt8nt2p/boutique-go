package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	pb "github.com/kx-boutique/api/auth/service/v1"
	userv1 "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/auth/service/internal/biz/password"
	"github.com/kx-boutique/app/auth/service/internal/conf"
	ent "github.com/kx-boutique/ent/generated"
	entModel "github.com/kx-boutique/ent/model"
	"github.com/kx-boutique/pkg/errors"
	"github.com/kx-boutique/pkg/model"
	"github.com/kx-boutique/pkg/util"
)

type AuthRepo interface {
	GetEntClient() *ent.Client
	Save(ctx context.Context, client *ent.Client, ae *entModel.Auth) *entModel.Auth
	FindPasswordHashByUserId(ctx context.Context, client *ent.Client, userId uuid.UUID) string
}

type AuthUseCase struct {
	appConfig  *conf.App
	userClient userv1.UserClient
	repo       AuthRepo
	log        *log.Helper
}

func NewAuthUseCase(appConfig *conf.App, userClient userv1.UserClient, repo AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		appConfig:  appConfig,
		userClient: userClient,
		repo:       repo,
		log:        log.NewHelper(log.With(logger, "module", "usecase/auth"))}
}

func (uc *AuthUseCase) NewAuth(ctx context.Context, req *pb.RegisterReq) *entModel.Auth {
	userId := util.ParseUUID(req.UserId)
	hash, err := password.GeneratePasswordHash(req.Password)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	entity := uc.repo.Save(ctx, uc.repo.GetEntClient(), &entModel.Auth{
		UserId:       userId,
		PasswordHash: hash,
	})

	return entity
}

func (uc *AuthUseCase) Login(ctx context.Context, req *pb.LoginReq) string {
	resp, err1 := uc.userClient.GetIdByEmail(ctx, &userv1.GetIdByEmailReq{
		Email: req.Email,
	})
	if err1 != nil {
		panic(errors.AppInternalErr(err1.Error()))
	}

	id := util.ParseUUID(resp.Id)
	passwordHash := uc.repo.FindPasswordHashByUserId(ctx, uc.repo.GetEntClient(), id)

	ok, _ := password.ComparePassword(req.Password, passwordHash)
	if !ok {
		panic(errors.AppInvalidCredentialsErr())
	}

	token, err2 := password.GenerateJWT(uc.appConfig.Jwt, id, req.Email)
	if err2 != nil {
		panic(errors.AppInvalidCredentialsErr())
	}

	return token
}

func (uc *AuthUseCase) ExtractJWTClaims(ctx context.Context) *model.Me {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		uc.log.Error("cannot extract claims from jwt")
		panic(errors.AppInvalidCredentialsErr())
	}

	mapClaims := claims.(jwtv4.MapClaims)
	return &model.Me{
		UserId: mapClaims["user_id"].(string),
		Email:  mapClaims["sub"].(string),
	}
}
