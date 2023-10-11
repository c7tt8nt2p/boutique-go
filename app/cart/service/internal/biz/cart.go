package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/kx-boutique/api/cart/service/v1"
	v1 "github.com/kx-boutique/api/product/service/v1"
	"github.com/kx-boutique/app/cart/service/internal/data"
	"github.com/kx-boutique/pkg/errors"
	"github.com/kx-boutique/pkg/util"
)

type CartUseCase struct {
	productClient v1.ProductClient

	cartRepo     data.CartRepo
	cartItemRepo data.CartItemRepo
	log          *log.Helper
}

func NewCartUseCase(productClient v1.ProductClient, cartRepo data.CartRepo, cartItemRepo data.CartItemRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{
		productClient: productClient,
		cartRepo:      cartRepo,
		cartItemRepo:  cartItemRepo,
		log:           log.NewHelper(log.With(logger, "module", "usecase/cart")),
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
	if err := validateAddItemToCart(ctx, uc, req); err != nil {
		return nil, err
	}

	myself := util.Myself(ctx)
	userId, err1 := util.ParseUUID(myself.UserId)
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

	if err4 := validateItemAlreadyInCart(ctx, uc, cartId, productId); err4 != nil {
		return nil, err4
	}

	return doAddItemToCart(ctx, uc, cartId, productId, req.Qty)
}

func validateAddItemToCart(ctx context.Context, uc *CartUseCase, req *pb.AddItemToCartReq) error {
	_, err := uc.productClient.ValidatePurchasableProduct(ctx, &v1.ValidatePurchasableProductReq{
		Id:  req.ProductId,
		Qty: req.Qty,
	})
	if err != nil {
		return err
	}
	return nil
}

func validateItemAlreadyInCart(ctx context.Context, uc *CartUseCase, cartId uuid.UUID, productId uuid.UUID) error {
	exist, err := uc.cartItemRepo.ExistsByCartIdAndProductId(ctx, cartId, productId)
	if err != nil {
		return err
	}
	if exist {
		return errors.ErrValidationFailed("Item already exists in the cart")
	}
	return nil
}

func doAddItemToCart(ctx context.Context, uc *CartUseCase, cartId uuid.UUID, productId uuid.UUID, qty int32) (*data.CartItemEntity, error) {
	return uc.cartItemRepo.Save(ctx, &data.CartItemEntity{
		CartId:    cartId,
		ProductId: productId,
		Qty:       qty,
	})
}

func (uc *CartUseCase) GetCart(ctx context.Context, uid int64) (*data.CartEntity, error) {
	return nil, nil
}
