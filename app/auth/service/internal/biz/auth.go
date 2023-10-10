package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/auth/service/v1"
	"github.com/kx-boutique/app/auth/service/internal/biz/password"
	"github.com/kx-boutique/app/auth/service/internal/data"
	"github.com/kx-boutique/pkg/util"
)

type AuthUseCase struct {
	repo data.AuthRepo
	log  *log.Helper
}

func NewAuthUseCase(repo data.AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/auth"))}
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

	entity, err3 := uc.repo.Save(ctx, &data.AuthEntity{
		UserId:       userId,
		PasswordHash: hash,
	})
	if err3 != nil {
		return nil, err3
	}

	return entity, nil
}
