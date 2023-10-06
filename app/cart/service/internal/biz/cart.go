package biz

import (
	"context"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/cart/service/internal/data"
	"github.com/kx-boutique/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type CartUseCase struct {
	repo data.CartRepo
	log  *log.Helper
}

func NewCartUseCase(repo data.CartRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/cart"))}
}

func (uc *CartUseCase) NewCart(ctx context.Context, userId string) (uuid.UUID, error) {
	id, err := util.ParseUUID(userId)
	if err != nil {
		return uuid.Nil, err
	}

	return uc.repo.SaveNewCart(ctx, id)
}

func (uc *CartUseCase) GetCart(ctx context.Context, uid int64) (*data.CartEntity, error) {
	return uc.repo.GetCart(ctx, uid)
}
