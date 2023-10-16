package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	cartv1 "github.com/kx-boutique/api/cart/service/v1"
	productv1 "github.com/kx-boutique/api/product/service/v1"
	entModel "github.com/kx-boutique/app/checkout/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/pkg/errors"
	"github.com/kx-boutique/pkg/util"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CheckoutRepo interface {
	GetEntClient() *ent.Client
	Save(ctx context.Context, client *ent.Client, co *entModel.Checkout) *ent.Checkout
}

type CheckoutItemRepo interface {
	GetEntClient() *ent.Client
	SaveAll(ctx context.Context, client *ent.Client, cois []*entModel.CheckoutItem) []*ent.CheckoutItem
}

type CheckoutUseCase struct {
	cartClient    cartv1.CartClient
	productClient productv1.ProductClient

	checkoutRepo     CheckoutRepo
	checkoutItemRepo CheckoutItemRepo
	log              *log.Helper
}

func NewCheckoutUseCase(cartClient cartv1.CartClient, productClient productv1.ProductClient,
	checkoutRepo CheckoutRepo, checkoutItemRepo CheckoutItemRepo, logger log.Logger) *CheckoutUseCase {
	return &CheckoutUseCase{
		cartClient:       cartClient,
		productClient:    productClient,
		checkoutRepo:     checkoutRepo,
		checkoutItemRepo: checkoutItemRepo,
		log:              log.NewHelper(log.With(logger, "module", "usecase/checkout"))}
}

func (uc *CheckoutUseCase) Checkout(ctx context.Context, _ *emptypb.Empty) *ent.Checkout {
	cart, err := uc.cartClient.ViewMyCart(ctx, &empty.Empty{})
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	cartItemIds := make([]string, len(cart.Items))
	// validate if products in the cart are still purchasable
	for i, ci := range cart.Items {
		cartItemIds[i] = ci.Id

		resp, err := uc.productClient.ValidatePurchasableProduct(ctx, &productv1.ValidatePurchasableProductReq{
			Id:  ci.ProductId,
			Qty: ci.Qty,
		})
		if err != nil {
			panic(errors.AppInternalErr(err.Error()))
		}
		if !resp.Validated {
			panic(errors.AppValidationErr(fmt.Sprintf("Error purchasing '%s': %s", ci.Name, resp.ValidationMessage)))
		}
	}

	txFunc := func(tx *ent.Tx) any {
		// save checkout and checkout items
		co := calculateCheckout(ctx, cart)
		com := uc.checkoutRepo.Save(ctx, uc.checkoutRepo.GetEntClient(), co)
		cois := calculateCheckoutItem(cart, com.ID)
		uc.checkoutItemRepo.SaveAll(ctx, uc.checkoutRepo.GetEntClient(), cois)

		// update product stocks
		for _, ci := range cart.Items {
			_, err := uc.productClient.SubtractProductStock(ctx, &productv1.SubtractProductStockReq{
				Id:            ci.ProductId,
				SubtractStock: ci.Qty,
			})
			if err != nil {
				panic(errors.AppInternalErr(err.Error()))
			}
		}

		// set checked out to true
		_, checkOutErr := uc.cartClient.CheckOutCartItem(ctx, &cartv1.CheckOutCartItemReq{
			Ids: cartItemIds,
		})
		if checkOutErr != nil {
			panic(errors.AppInternalErr(checkOutErr.Error()))
		}

		return com
	}
	return util.WithTx(ctx, uc.checkoutRepo.GetEntClient(), txFunc).(*ent.Checkout)
}

func calculateCheckout(ctx context.Context, cart *cartv1.ViewMyCartResp) *entModel.Checkout {
	totalPrice := decimal.NewFromFloat(0.0)
	for _, e := range cart.Items {
		totalPriceThisItem := decimal.NewFromFloat(e.Price).Mul(decimal.NewFromInt32(e.Qty))
		totalPrice = totalPrice.Add(totalPriceThisItem)
	}

	userId := util.ParseUUID(util.Me(ctx).UserId)
	return &entModel.Checkout{
		TotalPrice: totalPrice.InexactFloat64(),
		UserId:     userId,
	}
}

func calculateCheckoutItem(cart *cartv1.ViewMyCartResp, checkoutId uuid.UUID) []*entModel.CheckoutItem {
	var checkoutItems []*entModel.CheckoutItem
	for _, e := range cart.Items {
		id := util.ParseUUID(e.Id)
		checkoutItems = append(checkoutItems, &entModel.CheckoutItem{
			CartItemId: id,
			CheckoutId: checkoutId,
		})
	}
	return checkoutItems
}
