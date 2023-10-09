package biz

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/kx-boutique/api/cart/service/v1"
	"github.com/kx-boutique/app/cart/service/internal/data"
	"github.com/kx-boutique/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type CartUseCase struct {
	cartRepo     data.CartRepo
	cartItemRepo data.CartItemRepo
	log          *log.Helper
}

func NewCartUseCase(cartRepo data.CartRepo, cartItemRepo data.CartItemRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{
		cartRepo:     cartRepo,
		cartItemRepo: cartItemRepo,
		log:          log.NewHelper(log.With(logger, "module", "usecase/cart")),
	}
}

func (uc *CartUseCase) NewCart(ctx context.Context, req *pb.NewCartReq) (uuid.UUID, error) {
	id, err := util.ParseUUID(req.UserId)
	if err != nil {
		return uuid.Nil, err
	}

	return uc.cartRepo.Save(ctx, id)
}

func (uc *CartUseCase) AddItemToCart(ctx context.Context, req *pb.AddItemToCartReq) (*data.CartItemEntity, error) {
	userId, err1 := util.ParseUUID(req.UserId)
	if err1 != nil {
		return nil, err1
	}

	cartId, err2 := uc.cartRepo.FindIdByUserId(ctx, userId)
	if err2 != nil {
		return nil, err2
	}

	productId, err3 := util.ParseUUID(req.ProductId)
	if err3 != nil {
		return nil, err3
	}

	return uc.cartItemRepo.Save(ctx, &data.CartItemEntity{
		CartId:    cartId,
		ProductId: productId,
		Qty:       req.Qty,
	})
}

func (uc *CartUseCase) GetCart(ctx context.Context, uid int64) (*data.CartEntity, error) {
	return nil, nil
}
