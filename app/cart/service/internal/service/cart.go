package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/cart/service/internal/biz"

	pb "github.com/kx-boutique/api/cart/service/v1"
)

type CartService struct {
	pb.UnimplementedCartServer

	uc  *biz.CartUseCase
	log *log.Helper
}

func NewCartService(uc *biz.CartUseCase, logger log.Logger) *CartService {
	return &CartService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/cart"))}
}

func (s *CartService) NewCart(ctx context.Context, req *pb.NewCartReq) (*pb.NewCartResp, error) {
	id, err := s.uc.NewCart(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.NewCartResp{
		Id: id.String(),
	}, nil
}

func (s *CartService) ViewCart(ctx context.Context, req *pb.ViewCartReq) (*pb.ViewCartResp, error) {
	reply := &pb.ViewCartResp{Items: make([]*pb.ViewCartResp_Item, 0)}
	return reply, nil
}
