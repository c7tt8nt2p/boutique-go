package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	cartv1 "github.com/kx-boutique/api/cart/service/v1"
	pb "github.com/kx-boutique/api/checkout/service/v1"
	"github.com/kx-boutique/pkg/errors"
)

type CheckoutRepo interface {
}

type CheckoutUseCase struct {
	cartClient cartv1.CartClient

	repo CheckoutRepo
	log  *log.Helper
}

func NewCheckoutUseCase(cartClient cartv1.CartClient, repo CheckoutRepo, logger log.Logger) *CheckoutUseCase {
	return &CheckoutUseCase{cartClient: cartClient, repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/checkout"))}
}

func (uc *CheckoutUseCase) Checkout(ctx context.Context, req *pb.CheckoutReq) {
	cart, err := uc.cartClient.ViewCart(ctx, &empty.Empty{})
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}

	fmt.Println("my cart >>> ", cart)
}
