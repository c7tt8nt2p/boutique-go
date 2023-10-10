package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/auth/service/v1"
	userv1 "github.com/kx-boutique/api/user/service/v1"
	"github.com/kx-boutique/app/auth/service/internal/biz/password"
	"github.com/kx-boutique/app/auth/service/internal/data"
	"github.com/kx-boutique/pkg/util"
)

var (
	ErrInvalidCredentials = errors.Unauthorized("authentication failed", "invalid credentials")
)

type AuthUseCase struct {
	userClient userv1.UserClient
	repo       data.AuthRepo
	log        *log.Helper
}

func NewAuthUseCase(userClient userv1.UserClient, repo data.AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
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
		return "", ErrInvalidCredentials
	}

	id, err2 := util.ParseUUID(resp.Id)
	if err2 != nil {
		return "", ErrInvalidCredentials
	}

	passwordHash, err3 := uc.repo.FindPasswordHashByUserId(ctx, uc.repo.GetEntClient(), id)
	if err3 != nil {
		return "", ErrInvalidCredentials
	}

	ok, xxx := password.ComparePassword(req.Password, passwordHash)
	if !ok {
		fmt.Println("	invalid password", xxx)
		return "", ErrInvalidCredentials
	}
	fmt.Println("ok password")

	return passwordHash, nil
}
