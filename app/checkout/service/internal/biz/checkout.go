package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/checkout/service/v1"
)

type CheckoutRepo interface {

}

type CheckoutUseCase struct {
	repo CheckoutRepo
	log  *log.Helper
}

func NewCheckoutUseCase(repo CheckoutRepo, logger log.Logger) *CheckoutUseCase {
	return &CheckoutUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/checkout"))}
}

func (uc *CheckoutUseCase) Checkout(ctx context.Context, req *pb.CheckoutReq) {
}
