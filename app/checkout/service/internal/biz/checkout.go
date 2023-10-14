package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	cartv1 "github.com/kx-boutique/api/cart/service/v1"
	pb "github.com/kx-boutique/api/checkout/service/v1"
	entModel "github.com/kx-boutique/app/checkout/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/pkg/errors"
	"github.com/kx-boutique/pkg/util"
	"github.com/shopspring/decimal"
)

type CheckoutRepo interface {
	GetEntClient() *ent.Client
	Save(ctx context.Context, client *ent.Client, co *entModel.Checkout) uuid.UUID
}

type CheckoutItemRepo interface {
	GetEntClient() *ent.Client
	SaveAll(ctx context.Context, client *ent.Client, cois []*entModel.CheckoutItem) []*ent.CheckoutItem
}

type CheckoutUseCase struct {
	cartClient cartv1.CartClient

	checkoutRepo     CheckoutRepo
	checkoutItemRepo CheckoutItemRepo
	log              *log.Helper
}

func NewCheckoutUseCase(cartClient cartv1.CartClient, checkoutRepo CheckoutRepo, checkoutItemRepo CheckoutItemRepo, logger log.Logger) *CheckoutUseCase {
	return &CheckoutUseCase{
		cartClient:       cartClient,
		checkoutRepo:     checkoutRepo,
		checkoutItemRepo: checkoutItemRepo,
		log:              log.NewHelper(log.With(logger, "module", "usecase/checkout"))}
}

func (uc *CheckoutUseCase) Checkout(ctx context.Context, req *pb.CheckoutReq) {
	cart, err := uc.cartClient.ViewCart(ctx, &empty.Empty{})
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	txFunc := func(tx *ent.Tx) any {
		co := calculateCheckout(ctx, cart)
		coId := uc.checkoutRepo.Save(ctx, uc.checkoutRepo.GetEntClient(), co)
		cois := calculateCheckoutItem(cart, coId)
		return uc.checkoutItemRepo.SaveAll(ctx, uc.checkoutRepo.GetEntClient(), cois)
	}
	result := util.WithTx(ctx, uc.checkoutRepo.GetEntClient(), txFunc)

	fmt.Println("checkout >>> ", result)
}

func calculateCheckout(ctx context.Context, cart *cartv1.ViewCartResp) *entModel.Checkout {
	totalPrice := decimal.NewFromFloat(0.0)
	for _, e := range cart.Items {
		totalPriceThisItem := decimal.NewFromFloat(e.Price).Mul(decimal.NewFromInt32(e.Quantity))
		totalPrice = totalPrice.Add(totalPriceThisItem)
	}

	userId := util.ParseUUID(util.Me(ctx).UserId)
	return &entModel.Checkout{
		TotalPrice: totalPrice.InexactFloat64(),
		UserId:     userId,
	}
}

func calculateCheckoutItem(cart *cartv1.ViewCartResp, checkoutId uuid.UUID) []*entModel.CheckoutItem {
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
