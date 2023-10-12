package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/checkout/service/v1"
	"github.com/kx-boutique/app/checkout/service/internal/biz"
)

type CheckoutService struct {
	pb.UnimplementedCheckoutServer

	uc  *biz.CheckoutUseCase
	log *log.Helper
}

func NewCheckoutService(uc *biz.CheckoutUseCase, logger log.Logger) *CheckoutService {
	return &CheckoutService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/checkout")),
	}
}

func (s *CheckoutService) Checkout(ctx context.Context, req *pb.CheckoutReq) (*pb.CheckoutResp, error) {
	s.uc.Checkout(ctx, req)

	return &pb.CheckoutResp{}, nil
}
