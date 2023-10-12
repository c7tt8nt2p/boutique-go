package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/kx-boutique/api/cart/service/v1"
	v1 "github.com/kx-boutique/api/product/service/v1"
	entModel "github.com/kx-boutique/app/cart/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/pkg/errors"
	"github.com/kx-boutique/pkg/util"
)

type CartRepo interface {
	GetEntClient() *ent.Client
	Save(ctx context.Context, client *ent.Client, userId uuid.UUID) uuid.UUID
	FindIdByUserId(ctx context.Context, client *ent.Client, userId uuid.UUID) uuid.UUID
}

type CartItemRepo interface {
	GetEntClient() *ent.Client
	Save(ctx context.Context, client *ent.Client, cie *entModel.CartItem) *entModel.CartItem
	FindByCartId(ctx context.Context, client *ent.Client, cartId uuid.UUID) *entModel.CartWithProducts
	ExistsByCartIdAndProductId(ctx context.Context, client *ent.Client, cartId uuid.UUID, productId uuid.UUID) bool
	DeleteByCartIdAndProductId(ctx context.Context, client *ent.Client, cartId uuid.UUID, productId uuid.UUID) uuid.UUID
}

type CartUseCase struct {
	productClient v1.ProductClient

	cartRepo     CartRepo
	cartItemRepo CartItemRepo
	log          *log.Helper
}

func NewCartUseCase(productClient v1.ProductClient, cartRepo CartRepo, cartItemRepo CartItemRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{
		productClient: productClient,
		cartRepo:      cartRepo,
		cartItemRepo:  cartItemRepo,
		log:           log.NewHelper(log.With(logger, "module", "usecase/cart")),
	}
}

func findMyCartId(ctx context.Context, uc *CartUseCase) uuid.UUID {
	myself := util.Me(ctx)
	userId := util.ParseUUID(myself.UserId)
	cartId := uc.cartRepo.FindIdByUserId(ctx, uc.cartRepo.GetEntClient(), userId)
	return cartId
}

func (uc *CartUseCase) NewCart(ctx context.Context, req *pb.NewCartReq) uuid.UUID {
	id := util.ParseUUID(req.UserId)

	return uc.cartRepo.Save(ctx, uc.cartRepo.GetEntClient(), id)
}

func (uc *CartUseCase) AddItemToCart(ctx context.Context, req *pb.AddItemToCartReq) *entModel.CartItem {
	validateAddItemToCart(ctx, uc, req)

	cartId := findMyCartId(ctx, uc)
	productId := util.ParseUUID(req.ProductId)

	validateItemAlreadyInCart(ctx, uc, cartId, productId)

	return uc.cartItemRepo.Save(ctx, uc.cartItemRepo.GetEntClient(),
		&entModel.CartItem{
			CartId:    cartId,
			ProductId: productId,
			Qty:       req.Qty,
		})
}

func (uc *CartUseCase) RemoveItemFromCart(ctx context.Context, req *pb.RemoveItemFromCartReq) uuid.UUID {
	cartId := findMyCartId(ctx, uc)
	productId := util.ParseUUID(req.ProductId)

	return uc.cartItemRepo.DeleteByCartIdAndProductId(ctx, uc.cartItemRepo.GetEntClient(), cartId, productId)
}

func validateAddItemToCart(ctx context.Context, uc *CartUseCase, req *pb.AddItemToCartReq) {
	v, err := uc.productClient.ValidatePurchasableProduct(ctx, &v1.ValidatePurchasableProductReq{
		Id:  req.ProductId,
		Qty: req.Qty,
	})
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}
	if !v.Validated {
		panic(errors.AppInternalErr(v.ValidationMessage))
	}
}

func validateItemAlreadyInCart(ctx context.Context, uc *CartUseCase, cartId uuid.UUID, productId uuid.UUID) {
	exist := uc.cartItemRepo.ExistsByCartIdAndProductId(ctx, uc.cartItemRepo.GetEntClient(), cartId, productId)
	if exist {
		panic(errors.AppValidationErr("Item already exists in the cart"))
	}
}

func (uc *CartUseCase) ViewCart(ctx context.Context) *entModel.CartWithProducts {
	myself := util.Me(ctx)
	userId := util.ParseUUID(myself.UserId)

	cartId := uc.cartRepo.FindIdByUserId(ctx, uc.cartRepo.GetEntClient(), userId)
	result := uc.cartItemRepo.FindByCartId(ctx, uc.cartItemRepo.GetEntClient(), cartId)

	return result
}
