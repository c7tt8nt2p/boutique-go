package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/checkout/service/v1"
	"github.com/kx-boutique/app/checkout/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *CheckoutService) Checkout(ctx context.Context, _ *emptypb.Empty) (*pb.CheckoutResp, error) {
	co := s.uc.Checkout(ctx, &emptypb.Empty{})

	return &pb.CheckoutResp{
		Id:         co.ID.String(),
		TotalPrice: co.TotalPrice,
		CreatedAt:  timestamppb.New(co.CreatedAt),
	}, nil
}
