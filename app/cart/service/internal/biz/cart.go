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
	id := util.ParseUUID(req.UserId)

	return uc.cartRepo.Save(ctx, uc.cartRepo.GetEntClient(), id)
}

func (uc *CartUseCase) AddItemToCart(ctx context.Context, req *pb.AddItemToCartReq) (*data.CartItemEntity, error) {
	if err := validateAddItemToCart(ctx, uc, req); err != nil {
		return nil, err
	}

	myself := util.Myself(ctx)
	userId := util.ParseUUID(myself.UserId)

	cartId, err2 := uc.cartRepo.FindIdByUserId(ctx, uc.cartRepo.GetEntClient(), userId)
	if err2 != nil {
		return nil, err2
	}

	productId := util.ParseUUID(req.ProductId)

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
	exist, err := uc.cartItemRepo.ExistsByCartIdAndProductId(ctx, uc.cartItemRepo.GetEntClient(), cartId, productId)
	if err != nil {
		return err
	}
	if exist {
		panic(errors.AppValidationErr("Item already exists in the cart"))
	}
	return nil
}

func doAddItemToCart(ctx context.Context, uc *CartUseCase, cartId uuid.UUID, productId uuid.UUID, qty int32) (*data.CartItemEntity, error) {
	return uc.cartItemRepo.Save(ctx, uc.cartItemRepo.GetEntClient(),
		&data.CartItemEntity{
			CartId:    cartId,
			ProductId: productId,
			Qty:       qty,
		})
}

func (uc *CartUseCase) ViewCart(ctx context.Context) ([]*data.CartItemProductEntity, error) {
	myself := util.Myself(ctx)
	userId := util.ParseUUID(myself.UserId)

	cartId, err2 := uc.cartRepo.FindIdByUserId(ctx, uc.cartRepo.GetEntClient(), userId)
	if err2 != nil {
		return nil, err2
	}

	result, err3 := uc.cartItemRepo.FindByCartId(ctx, uc.cartItemRepo.GetEntClient(), cartId)
	if err3 != nil {
		return nil, err3
	}

	return result, nil
}
